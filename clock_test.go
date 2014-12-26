package clock

import (
	"testing"
	"time"
)

func TestRealClock(t *testing.T) {
	var (
		timeNow = time.Now()
		delta   = 1 * time.Millisecond
		now     = Now()
	)

	if timeNow.Round(delta) != now.Round(delta) {
		t.Error("Expected time.Now() to match Now()")
	}
}

func TestStubClock(t *testing.T) {
	tz, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	st := time.Date(1985, 10, 31, 10, 4, 0, 0, tz)
	Now = StubClock(st).Now

	time.Sleep(1 * time.Millisecond)

	if Now() != st {
		t.Error("Expected StubClock's Now() to always return same time")
	}
}
