package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}
	queue := make(chan Task, len(tasks))
	// Producer
	for _, task := range tasks {
		queue <- task
	}
	close(queue)
	// Consumer
	var wg sync.WaitGroup
	var once sync.Once
	var counterErrors int32
	var errorFlag bool
	wg.Add(n)
	for i := 0; i < n; i++ {
		//nolint:all
		//линтер жалуется на то, что всегда вывожу nil, но иначе горутину не завершить
		go func(wg *sync.WaitGroup) error {
			for task := range queue {
				if int(atomic.LoadInt32(&counterErrors)) >= m {
					once.Do(func() {
						errorFlag = true
					})
					wg.Done()
					return nil
				}
				if task() != nil {
					atomic.AddInt32(&counterErrors, 1)
				}
			}
			wg.Done()
			return nil
		}(&wg)
	}
	wg.Wait()
	if errorFlag {
		return ErrErrorsLimitExceeded
	}
	return nil
}
