package routing

import (
	"math/rand"
	"time"
)

type Stop struct {
	Err error
}

func (s Stop) Error() string {
	return s.Err.Error()
}

func Try(attempts int, sleep time.Duration, fn func() error) (err error) {
	if err = fn(); err != nil {
		if s, ok := err.(Stop); ok {
			return s.Err
		}

		if attempts--; attempts > 0 {
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2
			time.Sleep(sleep)
			return Try(attempts, 2*sleep, fn)
		}
	}
	return
}
