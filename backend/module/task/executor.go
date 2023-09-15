package task

import (
	"hidden-record-assistant/backend/module/zlog"
	"sync/atomic"
	"time"
)

type Task func()

type Executor struct {
	state uint32 // 0: running, 1: shutdown
	name  string
	ch    chan Task
}

//func NewExecutor(name string) *Executor {
//	return newExecutor(name, 20)
//}

func newExecutor(name string, chanCap int) *Executor {
	e := &Executor{
		state: 0,
		name:  name,
		ch:    make(chan Task, chanCap),
	}
	go e.loopRun()
	return e
}

func (e *Executor) loopRun() {
	defer func() {
		if err := recover(); err != nil {
			zlog.Errorf("executor [%s] panic: %v", e.name, err)
		}
	}()

	for task := range e.ch {
		if task == nil {
			if atomic.LoadUint32(&e.state) != 1 {
				zlog.Errorf("executor [%s] receive nil task, but state is not shutdown", e.name)
				break
			}
			break
		}
		task()
	}
}

func (e *Executor) Close() {
	if !atomic.CompareAndSwapUint32(&e.state, 0, 1) {
		zlog.Errorf("executor [%s] already shutdown", e.name)
		return
	}
	e.ch <- nil
}

func (e *Executor) Submit(name string, task Task) *Executor {
	if atomic.LoadUint32(&e.state) != 0 {
		zlog.Errorf("executor [%s] is shutdown, task [%s] Submit fail", e.name, name)
		return e
	}
	e.ch <- func() {
		start := time.Now()
		defer func() {
			zlog.Debugf("[%s] Cost: %s", name, time.Since(start).String())
		}()
		task()
	}
	return e
}
