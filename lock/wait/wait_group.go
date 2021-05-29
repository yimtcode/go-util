package wait

import (
	"sync"
	"time"
)

type WaitGroup interface {
	Wait(work func(), timeout time.Duration, keys ...interface{}) bool
	SetEvent(keys ...interface{})
}

func NewWaitGroup() WaitGroup {
	group := new(waitGroup)
	return group
}

type waitGroup struct {
	sync.Map
}

func (w *waitGroup) Wait(work func(), timeout time.Duration, keys ...interface{}) bool {
	// over delete
	defer func() {
		for _, key := range keys {
			w.Map.Delete(key)
		}
	}()

	// init lock
	for _, key := range keys {
		w.Map.Store(key, NewWait())
	}

	// start work
	go work()

	// wait finish
	wait := NewWait()
	over := false

	go func() {
		for _, key := range keys {
			if over {
				return
			}

			value, ok := w.Map.Load(key)
			if !ok {
				return
			}
			if !value.(Wait).Wait(timeout) {
				return
			}
		}

		wait.SetEvent()
	}()

	b := wait.Wait(timeout)
	over = true
	return b
}

func (w *waitGroup) SetEvent(keys ...interface{}) {
	for _, key := range keys {
		value, ok := w.Map.Load(key)
		if ok {
			value.(Wait).SetEvent()
		}
	}
}
