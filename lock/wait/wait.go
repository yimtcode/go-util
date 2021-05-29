package wait

import "time"

// Wait is interface
type Wait interface {
	Wait(timeout time.Duration) bool
	SetEvent()
}

func NewWait() Wait {
	w := new(wait)
	w.b = make(chan bool, 1)
	return w
}

type wait struct {
	b chan bool
}

func (w *wait) Wait(timeout time.Duration) bool {
	time.AfterFunc(timeout, func() {
		select {
		case w.b <- false:
		default:
		}
	})

	return <-w.b
}

func (w *wait) SetEvent() {
	select {
	case w.b <- true:
	default:
	}
}
