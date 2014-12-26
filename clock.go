package clock

import (
	"time"
)

// Now is a variable that holds a function that returns the current time.
// It can be swapped in runtime.
var Now func() time.Time

// initialize the default to RealNow
func init() {
	Now = RealNow
}

// RealNow returns the real currrent time and is the default
// value for Now.
func RealNow() time.Time {
	return time.Now()
}

// StubClock implements the Clock interface but it always returns
// the same time when Now() is called. It's most useful for testing.
type StubClock time.Time

// Now always returns the receiver, which is a time.
func (c StubClock) Now() time.Time {
	return time.Time(c)
}
