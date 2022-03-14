package microtime_test

import (
	"testing"
	"time"

	"github.com/KarpelesLab/microtime"
)

func must(t time.Time, e error) time.Time {
	if e != nil {
		panic(e)
	}
	return t
}

func TestTimeInt64(t *testing.T) {
	testV := []time.Time{
		must(time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")),
		must(time.Parse(time.RFC3339Nano, "1969-07-20T20:17:00.123456000Z")),
		must(time.Parse(time.RFC3339Nano, "0800-12-25T12:34:56.123456000Z")),
	}

	for _, v := range testV {
		intV := microtime.ToInt64(v)
		res := microtime.FromInt64(intV)

		// compare v == res
		if !v.Equal(res) {
			t.Errorf("int64 test failed: %s (%d) != %s", v, intV, res)
		}
	}
}

func TestTimeUint64(t *testing.T) {
	testV := []time.Time{
		must(time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")),
		must(time.Parse(time.RFC3339, "9999-07-01T12:34:56Z")),
		must(time.Parse(time.RFC3339Nano, "1970-01-01T00:00:04.123456000Z")),
	}

	for _, v := range testV {
		intV := microtime.ToUint64(v)
		res := microtime.FromUint64(intV)

		// compare v == res
		if !v.Equal(res) {
			t.Errorf("uint64 test failed: %s (%d) != %s", v, intV, res)
		}
	}
}
