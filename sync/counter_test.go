package sync

import (
	"sync"
	"testing"
)

type Counter struct {
	mutex sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
	})
	t.Run("it runs safely concurrently", func(t *testing.T) {
		var wg sync.WaitGroup
		wantedCount := 1000
		wg.Add(wantedCount)
		counter := Counter{}

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, &counter, wantedCount)
	})
}

func assertCounter(t *testing.T, counter *Counter, want int) {
	got := counter.Value()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
