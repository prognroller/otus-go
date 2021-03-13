package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if in == nil {
		return nil
	}

	for _, stage := range stages {
		tempIn := make(Bi)

		go func(in In) {
			defer close(tempIn)

			for {
				select {
				case v, ok := <-in:
					if !ok {
						return
					}

					tempIn <- v
				case <-done:
					return
				}
			}
		}(in)

		in = stage(tempIn)
	}

	return in
}
