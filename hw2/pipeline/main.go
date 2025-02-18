package main

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	curIn:=in
	for _,stage := range stages {	
		go func () {
			curIn=stage(curIn)
		}()
	} 
	return curIn
}


