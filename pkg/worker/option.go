package worker

func WithWorkerTimeout(timeout uint) func(wb *WorkerBase) {
	return func(wb *WorkerBase) {
		wb.workerTimeout = timeout
	}
}

func WithQueue(queue QueueItf) func(wb *WorkerBase) {
	return func(wb *WorkerBase) {
		wb.queue = queue
	}
}

func WithWorkerFunc(fn func([]byte) error) func(wb *WorkerBase) {
	return func(wb *WorkerBase) {
		wb.workerFunc = fn
	}
}
