package hw06pipelineexecution

import "sync"

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	startChannel := make(Bi)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer close(startChannel)
		defer wg.Done()
		for inElement := range in {
			select {
			case startChannel <- inElement:
			case <-done:
				return
			}
		}
	}()

	var inChannel In = startChannel
	var outChannel Out = nil
	for _, stage := range stages {
		outChannel = stage(inChannel)
		inChannel = outChannel
	}
	//wg.Wait()
	return outChannel
}
