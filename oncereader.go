package limit

import "sync"

type OnceReader[T any] struct {
	value T
	once  *sync.Once
}

func NewOnceReader[T any](v T) *OnceReader[T] {
	return &OnceReader[T]{v, new(sync.Once)}
}

func (r *OnceReader[T]) Read() (v T, err error) {
	onceReaded := false
	r.once.Do(func() { onceReaded = true })
	if onceReaded == false {
		return v, ReadLimitExceededError
	}
	return r.value, nil
}
