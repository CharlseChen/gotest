package goroutine_pool

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Task struct {
	Handler func() error
	Result  chan error
}

type Pool struct {
	capacity    int
	active      int64 // 使用 int64 以便与 atomic 包配合使用
	tasks       chan *Task
	quit        chan bool
	workerQueue chan *worker
	wg          sync.WaitGroup
}

type worker struct {
	pool *Pool
	task chan *Task
}

func NewPool(capacity int) *Pool {
	return &Pool{
		capacity:    capacity,
		tasks:       make(chan *Task),
		quit:        make(chan bool),
		workerQueue: make(chan *worker, capacity),
	}
}

func (p *Pool) Submit(task *Task) {
	select {
	case p.tasks <- task:
	case <-p.quit:
		return
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.capacity; i++ {
		w := &worker{
			pool: p,
			task: make(chan *Task),
		}
		go w.run()
	}

	go p.dispatch()
}

func (p *Pool) Stop() {
	close(p.quit)
	p.wg.Wait()
}

func (p *Pool) dispatch() {
	for {
		select {
		case task := <-p.tasks:
			worker := <-p.workerQueue
			worker.task <- task
		case <-p.quit:
			return
		}
	}
}

func (w *worker) run() {
	for {
		w.pool.workerQueue <- w

		select {
		case task := <-w.task:
			atomic.AddInt64(&w.pool.active, 1)
			w.pool.wg.Add(1)

			err := task.Handler()
			task.Result <- err

			w.pool.wg.Done()
			atomic.AddInt64(&w.pool.active, -1)

		case <-w.pool.quit:
			return
		}
	}
}

// ActiveWorkers 返回当前活跃的 worker 数量
func (p *Pool) ActiveWorkers() int {
	return int(atomic.LoadInt64(&p.active))
}

func main() {
	pool := NewPool(5)
	pool.Run()

	for i := 0; i < 10; i++ {
		taskID := i
		task := &Task{
			Handler: func() error {
				fmt.Printf("Task %d is running, active workers: %d\n", taskID, pool.ActiveWorkers())
				time.Sleep(time.Second)
				return nil
			},
			Result: make(chan error, 1),
		}

		pool.Submit(task)

		go func() {
			err := <-task.Result
			if err != nil {
				fmt.Printf("Task %d failed: %v\n", taskID, err)
			} else {
				fmt.Printf("Task %d completed successfully, active workers: %d\n", taskID, pool.ActiveWorkers())
			}
		}()
	}

	time.Sleep(3 * time.Second)
	pool.Stop()
	fmt.Println("Pool stopped, final active workers:", pool.ActiveWorkers())
}
