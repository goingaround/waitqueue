package waitqueue

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestWaitQueue_InvalidQueueSize(t *testing.T) {
	_, err := New(0)
	require.Error(t, err)
}

func TestWaitQueue_OK(t *testing.T) {
	wq, err := New(1)
	require.NoError(t, err)

	counter := 0

	enq := func() {
		go func() {
			wq.Enq()
			counter++
		}()
		time.Sleep(5*time.Millisecond) // give the goroutine a chance to do it's work
	}

	deq := func() {
		wq.Deq()
		counter--
		time.Sleep(5*time.Millisecond) // give waiting enq a chance to do it's work
	}

	// starting off with zero
	require.Equal(t, 0, counter)
	// enq once and counter should be 1
	enq()
	require.Equal(t, 1, counter)
	// enq once again, ad queue size reached this one will cause a wait
	// counter should still be 1
	enq()
	require.Equal(t, 1, counter)
	// deq will decrease the counter to 0 but prev enq waiting for the contention
	// will be released and will increase counter to 1 again
	deq()
	require.Equal(t, 1, counter)
	// now no enq is waiting, deq will decrease the counter to 0
	deq()
	require.Equal(t, 0, counter)
}
