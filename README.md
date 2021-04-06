_WaitQueue_ is an object based on _sync.WaitGroup_ added with a buffered channel which acts as a queue contention locker.

`New(size int) (*WaitQueue, error)` - initializes a new wait queue

`(g *WaitQueue) Enq()` - Add(1) on inner wait group

`(g *WaitQueue) Deq()` - Done() on inner wait group

`func (g *WaitQueue) Wait()` - Wait() on inner wait group