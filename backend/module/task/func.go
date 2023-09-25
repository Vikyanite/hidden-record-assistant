package task

func ExecuteTaskConcurrently(tasks ...func() error) (errChan chan error, num int) {
	num = len(tasks)
	errChan = make(chan error, num)
	for i := range tasks {
		go func(index int) {
			errChan <- tasks[index]()
		}(i)
	}
	return
}
