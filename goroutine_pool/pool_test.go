package goroutine_pool

import (
	"testing"
	"time"
	"fmt"
)

func TestPool(t *testing.T) {
	pool := NewPool(10)
	pool.Start()
	defer pool.Stop()

	for i := 0; i < 1000; i++ {
		s := i
		task := &Task{
			Handler: func() error {
				time.Sleep(time.Duration(s) * time.Millisecond)
				fmt.Printf("Task %d completed\n", s)
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
	pool.Shutdown()
}
