package task

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

type ExecutorPool struct {
	index, size uint32
	executors   []*Executor
}

func NewExecutorPool(namePre string) *ExecutorPool {
	num := runtime.NumCPU()
	executors := make([]*Executor, num)
	for i := 0; i < num; i++ {
		executors[i] = newExecutor(fmt.Sprintf("%s_%d", namePre, i), 20)
	}
	return &ExecutorPool{
		index:     0,
		size:      uint32(num),
		executors: executors,
	}
}

func (e *ExecutorPool) Fixed(index uint32) *Executor {
	return e.executors[index%e.size]
}

func (e *ExecutorPool) Next() *Executor {
	i := atomic.AddUint32(&e.index, 1) % e.size
	return e.executors[i]
}

func (e *ExecutorPool) Close() {
	for _, executor := range e.executors {
		executor.Close()
	}
}
