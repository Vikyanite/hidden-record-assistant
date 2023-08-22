package zlog

import "log"

func MustNil(err error) {
	if err != nil {
		Fatalf("%v", err)
	}
}

func Recover() {
	if err := recover(); err != nil {
		log.Panicf("%v", err)
	}
}
