package microtime

import (
	"encoding/binary"
	"time"
)

// ToUint64 will return a uint64 version of the passed time. It is assumed
// time will always be after 1970, and no check is performed to see if it might
// actually be in the past. If you need to handle values in the past, use
// ToInt64.
func ToUint64(t time.Time) uint64 {
	// this method of encoding time will work until year 586524
	tm := uint64(t.Unix()) * 1000000
	tm += uint64(t.Nanosecond()) / 1000 // convert to microsecond

	return tm
}

// ToInt64 will return a int64 version of the passed time, with negative values
// corresponding to anything prior to January 1st 1970 UTC. This can be useful
// if working with an order dataset (in order to get a proper ordering,
// offsetting the set's start date will be needed).
func ToInt64(t time.Time) int64 {
	// this method of encoding time will work until year 294247
	return t.UnixMicro()
}

// ToBytes converts time to a binary value suitable for sorting. This is made
// to work with dates >=1970-01-01 and will not work properly for data prior to
// that.
func ToBytes(t time.Time) []byte {
	v := make([]byte, 8)
	binary.BigEndian.PutUint64(v, ToUint64(t))
	return v
}

// FromUint64 converts a given uint64 value back to time.Time.
func FromUint64(tm uint64) time.Time {
	ux := tm / 1000000
	return time.Unix(int64(ux), int64(tm-(ux*1000000))*1000)
}

// FromInt64 converts a given int64 value back to time.Time.
func FromInt64(tm int64) time.Time {
	return time.UnixMicro(tm)
}

// FromBytes converts a given binary array value to time.Time, expecting to
// only find values after 1970-01-01.
func FromBytes(v []byte) time.Time {
	return FromUint64(binary.BigEndian.Uint64(v))
}
