package waitqueue

import (
	"errors"
	"sync"
)

type WaitQueue struct {
	q  chan struct{}
	wg sync.WaitGroup
}

func New(size int) (*WaitQueue, error) {
	if size < 1 {
		return nil, errors.New("queue size must be bigger than 0")
	}
	return &WaitQueue{q: make(chan struct{}, size)}, nil
}

func (wq *WaitQueue) Enq() {
	wq.wg.Add(1)
	wq.q <- struct{}{}
}

func (wq *WaitQueue) Wait() {
	wq.wg.Wait()
}

func (wq *WaitQueue) Deq() {
	wq.wg.Done()
	<-wq.q
}
