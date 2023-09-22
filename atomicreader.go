package limit

import "sync/atomic"

type AtomicReader[T any] struct {
	value     T
	readCount int32
}

func NewAtomicReader[T any](v T, limit int32) *AtomicReader[T] {
	return &AtomicReader[T]{v, limit}
}

func (r *AtomicReader[T]) Read() (v T, err error) {
	if r.readCount <= 0 {
		return v, ReadLimitExceededError
	}
	atomic.AddInt32(&r.readCount, -1)
	return r.value, nil
}
