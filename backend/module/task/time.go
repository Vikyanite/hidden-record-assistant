package task

import (
	"hidden-record-assistant/backend/module/zlog"
	"time"
)

func TimeCost(name string) func() {
	start := time.Now()
	return func() {
		zlog.Debugf("[%s] Cost: %s", name, time.Since(start).String())
	}
}
