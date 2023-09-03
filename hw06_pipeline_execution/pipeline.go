package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func doneChannelWrap(in In, done In) Out {
	ch := make(Bi)
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case item, ok := <-in:
				if !ok {
					return
				}
				ch <- item
			}
		}
	}()
	return ch
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	input := in
	var output In
	for _, stage := range stages {
		output = stage(doneChannelWrap(input, done))
		input = output
	}
	return output
}
