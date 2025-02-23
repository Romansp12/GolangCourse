package main

import (
	"fmt"
	"time"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	curIn:=in
	for _,stage := range stages {
		// перед каждой стадией поставим фильтр который перед отправкой, будет проверять не отменено ли исполенение
		// CurIn filtered to res
		res:=make(chan interface{})
		go func(curChan In){
			for{
				select {
				case <-done:
					close(res)
					return
				default:
					x,clse:=<-curChan
					if !clse{
						close(res)
						return
					}
					select{
					case <-done:
						close(res)
						return
					default:
						res<-x
					}
				}
			}
		}(curIn)
		curIn=res
		curIn=stage(curIn)
	} 
	return curIn
}

func main() {
	stages:=make([]Stage,2)
	stages[0]=func(in In) (out Out) {
		res:=make(chan interface{})
		go func(){
			x:=<-in
			fmt.Println("stages 0",x)
			res<-x
		}()
		return res
	}

	stages[1]=func(in In) (out Out) {
		res:=make(chan interface{})
		go func(){
			x:=<-in
			fmt.Println("stages 1",x)
			res<-x
		}()
		return res
	}

	input:=make(chan interface{})
	done:=make(chan interface{})
	output:=ExecutePipeline(input,done,stages...)
	input<-123
	fmt.Println(<-output)
	time.Sleep(time.Second)
}


