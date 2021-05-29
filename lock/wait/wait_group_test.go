package wait

import (
	"testing"
	"time"
)

func TestWaitGroup_Wait(t *testing.T) {
	wg := NewWaitGroup()
	keyLogin1 := 1
	keyLogin2 := 2

	b := wg.Wait(func() {
		// login 1
		<-time.After(2 * time.Second)
		// login 2
		<-time.After(1 * time.Second)

		wg.SetEvent(keyLogin1, keyLogin2)
	}, 10*time.Second, keyLogin1, keyLogin2)

	if !b {
		t.Fatal("WaitGroup test failed")
	}
}
