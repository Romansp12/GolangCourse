package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)


func or(channels ...<- chan interface{}) <-chan struct{} {
	res:=make(chan struct{})
	var isClosed int32
	ctx,cancel:=context.WithCancel(context.Background())

	for _,ch:= range channels{
		go func(){
			select{			
			case <-ctx.Done():
				return
			case <-ch:
			}
			if atomic.CompareAndSwapInt32(&isClosed,0,1){
				cancel()
				close(res)
			}
		}()
	}
	return res
}

func main() {
	sig := func(after time.Duration) <- chan interface{} {
		c := make(chan interface{})
			go func() {
				defer close(c)
				time.Sleep(after)
		}()
		return c
	}
	
	start := time.Now()
	<-or (
		sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Second),
	sig(1*time.Second),
	sig(1*time.Second),
	sig(1*time.Second),
	sig(1*time.Second),
	sig(1*time.Second),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))
}