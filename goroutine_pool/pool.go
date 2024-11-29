package goroutine_pool

import (
	"fmt"
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
}

type worker struct {
	pool *Pool
	task chan *Task
	id   uint32
}

func NewPool(capacity int) *Pool {
	return &Pool{
		capacity:    capacity,
		tasks:       make(chan *Task, capacity*2),
		quit:        make(chan bool),
		workerQueue: make(chan *worker, capacity),
	}
}

func (p *Pool) Submit(task *Task) {
	//1.任务来了先进全局的任务channel排队
	p.tasks <- task
}

func (p *Pool) Run() {
	for i := 0; i < p.capacity; i++ {
		//一个工作协程承载一个任务，工作协程的数量就是池子的大小，也就是工作队列的大小
		w := &worker{
			pool: p,
			task: make(chan *Task),
			id:   uint32(i) + 1,
		}
		//将工作协程放入池中
		w.pool.workerQueue <- w
		go w.run()
	}

	go p.dispatch()
}

func (p *Pool) Stop() {
	close(p.quit)
}

func (p *Pool) dispatch() {
	//监听早在池子初始化的时候就开始了
	//2.在这里监听，读取放入，从工作队列中取出一个工作协程，承载任务，而具体承载任务的是worker中的task  channel
	for {
		select {
		case task := <-p.tasks:
			wor := <-p.workerQueue
			wor.task <- task
		case <-p.quit:
			return
		}
	}
}

func (w *worker) run() {
	//监听早在池子初始化的时候就开始了
	//这里监听的是工作协程的channel是否有任务
	for {
		select {
		case task := <-w.task:
			fmt.Printf("worker %d start\n", w.id)
			atomic.AddInt64(&w.pool.active, 1)
			//w.pool.wg.Add(1)

			err := task.Handler()
			task.Result <- err

			//w.pool.wg.Done()
			atomic.AddInt64(&w.pool.active, -1)
			//工作协程用完了,要还回去
			w.pool.workerQueue <- w
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
