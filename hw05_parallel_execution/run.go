package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, runnersLimit, errorsLimit int) error {
	var errors int32
	wg := sync.WaitGroup{}
	ch := tasksToChannel(tasks)

	for runner := 0; runner < runnersLimit; runner++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			startRunner(ch, &errors, int32(errorsLimit))
		}()
	}

	wg.Wait()
	if errors >= int32(errorsLimit) {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func tasksToChannel(tasks []Task) chan Task {
	ch := make(chan Task, len(tasks))
	defer close(ch)
	for _, task := range tasks {
		ch <- task
	}
	return ch
}

func startRunner(сhannel chan Task, errors *int32, errorsLimit int32) {
	for task := range *channel {
		if atomic.LoadInt32(errors) >= errorsLimit {
			return
		}
		err := task()
		if err != nil {
			atomic.AddInt32(errors, 1)
		}
	}
}
