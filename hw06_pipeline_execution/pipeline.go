package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	var connector Out
	for i := 0; i < len(stages); i++ {
		if i == 0 {
			connector = getNextChan(stages[i](in), done)
		} else {
			connector = getNextChan(stages[i](connector), done)
		}
	}
	return connector
}

func getNextChan(in, done In) Out {
	out := make(Bi)
	go func(out Bi) {
		defer close(out)
	Loop:
		for v := range in {
			select {
			case <-done:
				break Loop
			default:
				out <- v
			}
		}
	}(out)
	return out
}
