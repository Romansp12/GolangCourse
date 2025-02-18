package main

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"math/rand"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	ctx,cancel := context.WithCancel(context.Background())
	
	cntTask:=len(tasks)
	
	outputTask:=make(chan error,n)
	input:=make(chan Task)
	resultErr:=make(chan error,1)

	go func(){
		for e :=range outputTask{
			fmt.Println("receved error",e)
			if e!=nil{
				if m>=1{
					m--
				}else{
					cancel()
					resultErr<-ErrErrorsLimitExceeded
					close(resultErr)
					return
				}
			}
			cntTask--
			if cntTask==0{
				cancel()
				resultErr<-nil
				close(resultErr)
				return
			}
		}
	}()
	
	for i:=0;i<n;i++{
		go func(){
			for{
				select{
				case <-ctx.Done():
					return 
				default:
					task,closed:=<-input
					if !closed{
						return
					}
					select{
					case <-ctx.Done():
						return
					default:
						outputTask<-task()
					}
				}
			}
		}()
	}


	go func(){
		for _, task := range tasks {
			select{
			case <-ctx.Done():
				return
			default:
				select{
				case <-ctx.Done():
					return
				case input<-task:
				}
			}
		}
		close(input)
	}()

	return <-resultErr
}

func main() {
  tasksCount := 50
  tasks := make([]Task, 0, tasksCount)

  var runTasksCount int32

  for i := 0; i < tasksCount; i++ {
   err := fmt.Errorf("error from task %d", i)
   tasks = append(tasks, func() error {
	fmt.Printf("task start %d \n",i)
    time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
    atomic.AddInt32(&runTasksCount, 1)
    return err
   })
  }

  workersCount := 10
  maxErrorsCount := 23
  err := Run(tasks, workersCount, maxErrorsCount)
  fmt.Println(runTasksCount)
  fmt.Println(err)
}
