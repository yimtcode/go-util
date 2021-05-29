package wait

import (
	"testing"
	"time"
)

func TestWait_Wait(t *testing.T) {
	w := NewWait()

	time.AfterFunc(2*time.Second, func() {
		w.SetEvent()
	})

	b := w.Wait(10 * time.Second)
	if !b {
		t.Fatal("test wait failed")
	}
}
