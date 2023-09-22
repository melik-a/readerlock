package limit

import (
	"sync"
	"testing"
)

func TestReaders(t *testing.T) {
	readTests := []struct {
		name   string
		reader Reader[int]
		limit  int
		value  int
	}{
		{name: "ChannelReader", reader: NewChannelReader(10, 5), limit: 5, value: 10},
		{name: "AtomicReader", reader: NewAtomicReader(10, 3), limit: 3, value: 10},
		{name: "MutexReader", reader: NewMutexReader(10, 4), limit: 4, value: 10},
		{name: "OnceReader", reader: NewOnceReader(10), limit: 1, value: 10},
	}

	for _, test := range readTests {
		t.Run(test.name+" Read should return value when limit is not exceeded", func(t *testing.T) {
			assertReadAtLeastLimitTimes(t, test.reader, test.value, test.limit)
		})
		t.Run(test.name+" Read should not return value when limit is exceeded", func(t *testing.T) {
			_, gotErr := test.reader.Read()
			assertErrorWhenLimitExceeded(t, gotErr, ReadLimitExceededError)
		})
	}
}

func assertResult(t *testing.T, got, want interface{}) {
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertErrorWhenLimitExceeded(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("did not get error when expected")
	}
	assertResult(t, got, want)
}

func assertReadAtLeastLimitTimes[T comparable](t *testing.T, reader Reader[T], wantValue T, limit int) {
	t.Helper()
	var wg sync.WaitGroup
	var results chan T = make(chan T)
	for i := 0; i < limit; i++ {
		wg.Add(1)
		go func(r Reader[T], waitGroup *sync.WaitGroup) {
			defer wg.Done()
			result, err := r.Read()
			if err == nil {
				results <- result
			}
		}(reader, &wg)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	for value := range results {
		assertResult(t, value, wantValue)
	}
}
