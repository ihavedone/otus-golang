package hw05parallelexecution

import (
	"errors"
	"math"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, runnersLimit, errorsLimit int) error {
	var errors int32
	wg := sync.WaitGroup{}
	chankSize := len(tasks) / runnersLimit

	for runner := 0; runner < runnersLimit; runner++ {
		wg.Add(1)
		sliceFrom := runner * chankSize
		sliceTo := int(math.Min(float64(sliceFrom+chankSize), float64(len(tasks))))
		go func(tasksChank []Task) {
			defer wg.Done()
			startRunner(tasksChank, &errors, int32(errorsLimit))
		}(tasks[sliceFrom:sliceTo])
	}

	wg.Wait()
	if errors >= int32(errorsLimit) {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func startRunner(tasks []Task, errors *int32, errorsLimit int32) {
	for _, task := range tasks {
		if atomic.LoadInt32(errors) >= errorsLimit {
			return
		}
		err := task()
		if err != nil {
			atomic.AddInt32(errors, 1)
		}
	}
}
