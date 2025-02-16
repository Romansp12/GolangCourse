package main

import (
	"context"
	"errors"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Launch(n,m int) (input chan Task, resultErr chan error) {
	ctx,cancel := context.WithCancel(context.Background())
	
	outputTask:=make(chan error,n)
	input=make(chan Task,n)
	resultErr=make(chan error,1)
	
	go func(){
		for e :=range outputTask{
			if e!=nil{
				if m>=1{
					m--
				}else{
					cancel()
					resultErr<-ErrErrorsLimitExceeded
				}
			}
			n--
			if n==0{
				resultErr<-nil
			}
		}
	}()
	
	for i:=0;i<n;i++{
		go func(){
			for{
				select{
				case <-ctx.Done():
					return
				case task:=<-input:
					outputTask<-task()
				}
			}
		}()
	}
	
	return input,resultErr
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	input,resultErr := Launch(n,m)
	for _, task := range tasks {
		input<-task
	}
	return <-resultErr
}
