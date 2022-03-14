package microtime

import (
	"sync/atomic"
	"time"
)

var lastStampValue uint64

// Stamp returns a uint64 timestamp in microsecond that is unique and ever
// increasing, so that even if called multiple times in the same millisecond
// each call will return a different and increasing value.
//
// This can be safely called from multiple threads, and does not lock.
func Stamp() uint64 {
	tm := ToUint64(time.Now())

	v := atomic.LoadUint64(&lastStampValue)
	if v >= tm {
		// we are generating too many values or timestamp went back,
		// either way we can't use the timestamp, so increase
		// lastStampValue instead
		return atomic.AddUint64(&lastStampValue, 1) // returns lastStampValue+1
	}

	// store timestamp value since it was higher than the current value
	if atomic.CompareAndSwapUint64(&lastStampValue, v, tm) {
		// swap successful, now let's just return tm
		return tm
	}

	// compare and swap failed, means lastStampValue was updated between load and CaS, let's just return value+1
	return atomic.AddUint64(&lastStampValue, 1)
}
