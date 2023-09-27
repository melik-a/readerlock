package limit

type BufferedChannelReader[T any] struct {
	value       T
	syncChannel chan struct{}
}

func NewBufferedChannelReader[T any](v T, limit int) *BufferedChannelReader[T] {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	close(ch)
	return &BufferedChannelReader[T]{v, ch}
}

func (r *BufferedChannelReader[T]) Read() (v T, err error) {
	_, ok := <-r.syncChannel
	if !ok {
		return v, ReadLimitExceededError
	}
	return r.value, nil
}
