package goroutine_pool

import (
	"sync"
	"time"
)

type Task struct {
	Handler func() error
	Result  chan error
}

type Pool struct {
	capacity    int
	active      int
	tasks       chan *Task
	quit        chan bool
	workerQueue chan *worker
	mutex       sync.Mutex
}

type worker struct {
	pool *Pool
}

func NewPool(capacity int) *Pool {
	if capacity < 1 {
		capacity = 1
	}
	return &Pool{
		capacity:    capacity,
		tasks:       make(chan *Task, capacity*2),
		quit:        make(chan bool),
		workerQueue: make(chan *worker, capacity),
	}
}

func (p *Pool) Start() {
	for i := 0; i < p.capacity; i++ {
		w := &worker{pool: p}
		p.workerQueue <- w
		p.active++
		go w.run()
	}
}

func (w *worker) run() {
	for {
		select {
		case task := <-w.pool.tasks:
			if err := task.Handler(); err != nil {
				task.Result <- err
			} else {
				task.Result <- nil
			}
			w.pool.workerQueue <- w

		case <-w.pool.quit:
			return
		}
	}
}

func (p *Pool) Submit(task *Task) {
	p.tasks <- task
}

func (p *Pool) Stop() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.active > 0 {
		close(p.quit)
		p.active = 0
	}
}

func (p *Pool) Shutdown() {
	close(p.quit)
	close(p.tasks)
}

func (p *Pool) adjustWorkers() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		p.mutex.Lock()
		taskCount := len(p.tasks)
		workerCount := len(p.workerQueue)

		switch {
		case taskCount > workerCount && p.active < p.capacity:
			w := &worker{pool: p}
			p.workerQueue <- w
			p.active++
			go w.run()
		case taskCount < workerCount/2 && p.active > p.capacity/2:
			select {
			case w := <-p.workerQueue:
				p.active--
				w.pool.quit <- true
			default:
			}
		}
		p.mutex.Unlock()
	}
}

type BatchPool struct {
	*Pool
	batchSize int
	batchChan chan []*Task
}

func (bp *BatchPool) processBatch() {
	batch := make([]*Task, 0, bp.batchSize)
	timer := time.NewTimer(100 * time.Millisecond)

	for {
		select {
		case task := <-bp.tasks:
			batch = append(batch, task)
			if len(batch) == bp.batchSize {
				bp.batchChan <- batch
				batch = make([]*Task, 0, bp.batchSize)
			}
		case <-timer.C:
			if len(batch) > 0 {
				if len(batch) > 0 {
					bp.batchChan <- batch
					batch = make([]*Task, 0, bp.batchSize)
				}
			}
			timer.Reset(100 * time.Millisecond)
		}
	}
}
