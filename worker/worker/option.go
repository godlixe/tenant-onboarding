package worker

func WithWorkerNum(num uint) func(wb *WorkerBase) {
	return func(wb *WorkerBase) {
		wb.WorkerNum = num
	}
}

func WithWorkerTimeout(timeout uint) func(wb *WorkerBase) {
	return func(wb *WorkerBase) {
		wb.WorkerTimeout = timeout
	}
}

func WithQueue(queue QueueItf) func(wb *WorkerBase) {
	return func(wb *WorkerBase) {
		wb.queue = queue
	}
}

func WithWorkerFunc(fn func([]byte)) func(wb *WorkerBase) {
	return func(wb *WorkerBase) {
		wb.workerFunc = fn
	}
}
