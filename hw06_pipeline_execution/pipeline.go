package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// этот пайп для передачи канала от одного стейджа к другому
	readyChannels := make(chan Out, 1)
	// этот пайп выводится если сработал done
	terminateChannel := make(Bi)
	// заранее закрываем пайпы
	defer close(readyChannels)
	defer close(terminateChannel)
	for _, stage := range stages {
		select {
		case pipe := <-readyChannels:
			readyChannels <- stage(pipe)
		case <-done:
			return terminateChannel
		// здесь будет первое вхождение
		default:
			readyChannels <- stage(in)
		}
	}
	// после данного цикла все пайпы были прокинуты в нужные стейджи
	// перестрахуемся и сделаем еще один select(на тот случай если done все же пришел под конец)
	// это имеет смысл делать
	select {
	case <-done:
		return terminateChannel
	case output := <-readyChannels:
		return output
	}
}
