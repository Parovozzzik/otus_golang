package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in
	for _, stage := range stages {
		curIn := stage(out)
		curOut := make(Bi)

		go func() {
			defer close(curOut)

			for {
				select {
				case <-done:
					return
				case v, ok := <-curIn:
					if !ok {
						return
					}
					curOut <- v
				}
			}
		}()
		out = curOut
	}
	return out
}
