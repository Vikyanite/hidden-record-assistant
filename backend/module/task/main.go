package task

var DefaultTaskPool = NewExecutorPool("DTP")

func GetExecutor() *Executor {
	return DefaultTaskPool.Next()
}

func Submit(name string, task Task) *Executor {
	return DefaultTaskPool.Next().Submit(name, task)
}
