package limit

type ChannelReader[T any] struct {
	value       T
	readCount   int
	syncChannel chan struct{}
}

func NewChannelReader[T any](v T, limit int) *ChannelReader[T] {
	return &ChannelReader[T]{v, limit, make(chan struct{}, 1)}
}

func (r *ChannelReader[T]) Read() (v T, err error) {
	r.syncChannel <- struct{}{}
	if r.readCount <= 0 {
		return v, ReadLimitExceededError
	}
	r.readCount -= 1
	<-r.syncChannel
	return r.value, nil
}
