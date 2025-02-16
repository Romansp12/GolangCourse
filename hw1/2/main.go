package main

import (
	"context"
)


type Semaphore interface{
	Acquire(context.Context, int64) error
	TryAcquire(int64) bool
	Release(int64)
}


type ChanSemaphore struct{
	ch chan struct{}
}

func NewChanSemaphore(n int) *ChanSemaphore{
	return &ChanSemaphore{ch: make(chan struct{},n)}
}

func (s *ChanSemaphore) Acquire(ctx context.Context,cnt int64) {
	
} 