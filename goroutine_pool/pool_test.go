package goroutine_pool

import (
	"testing"
	"time"
	"fmt"
)

func TestPool(t *testing.T) {
	pool := NewPool(10)
	pool.Run()

	for i := 0; i < 100; i++ {
		s := i
		task := &Task{
			Handler: func() error {
				fmt.Printf("Task %d completed\n active workers:%d", s, pool.ActiveWorkers())
				return nil
			},
			Result: make(chan error, 1),
		}
		pool.Submit(task)
		go func(t *Task) {
			err := <-task.Result
			if err != nil {
				fmt.Printf("Task %d failed\n", s)
			}
		}(task)
	}

	time.Sleep(time.Second * 10)
	pool.Stop()
}
