package routing

import "sync/atomic"

type Lock struct {
	seed int32
}

func (l *Lock) Lock() bool {
	return atomic.CompareAndSwapInt32(&l.seed, 0, 1)
}

func (l *Lock) Unlock() {
	atomic.StoreInt32(&l.seed, 0)
}
