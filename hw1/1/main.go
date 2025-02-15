package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func calculate(i int,n int,res chan float64) {
	k:=i
	for {
		res<-(float64(2.0)/(float64(4*k+1)*float64(4*k+3)))
		k+=n
	}
}

func main() {
	var n int
	flag.IntVar(&n,"n", 2, "count gorutines")
	flag.Parse()

	if n<1 {
		fmt.Println("Invalid gorutines count")
	}

	fmt.Println("Running gorutines for counting: ",n)
	
	res:=make(chan float64,n)
	var ans float64
	go func(){
		for x:= range res{
			ans +=x
		}
	}()
	
	for i:=0;i<n;i++ {
		go calculate(i,n,res)
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
	fmt.Println(ans*4)
}