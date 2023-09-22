package limit

import (
	"sync"
)

type MutexSyncReader[T any] struct {
	mutex     *sync.Mutex
	value     T
	readCount int
}

func NewMutexReader[T any](v T, limit int) *MutexSyncReader[T] {
	return &MutexSyncReader[T]{new(sync.Mutex), v, limit}
}

func (r *MutexSyncReader[T]) Read() (v T, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if r.readCount <= 0 {
		return v, ReadLimitExceededError
	}
	r.readCount -= 1
	return r.value, nil
}
