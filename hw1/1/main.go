package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func calculate(ctx context.Context,i int,n int,res chan float64) {
	k:=i
	for {
		select {
		case <-ctx.Done():
			return
		case res<-(float64(2.0)/(float64(4*k+1)*float64(4*k+3))):
			k+=n
		}
	}
}

func main() {
	var n int
	flag.IntVar(&n,"n", 2, "count gorutines")
	flag.Parse()

	if n<1 {
		fmt.Println("Invalid gorutines count")
		return
	}
	
	ctx,cancel:=context.WithCancel(context.Background())
	res:=make(chan float64)
	var ans float64
	
	wg:=sync.WaitGroup{}
	wg.Add(1)
	go func(){
		defer wg.Done()
		for {
		select{
			case <-ctx.Done():
				return
			default:
				x,closed:=<-res
				if !closed{
					return
				}
				select{
				case <-ctx.Done():
					return
				default:
					ans+=x
				}
		}
		}
	}()
	
	for i:=0;i<n;i++ {
		go calculate(ctx,i,n,res)
	}
	fmt.Println("Running gorutines for counting: ",n)
	
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
	
	cancel()
	wg.Wait()
	
	fmt.Println(ans*4)
}