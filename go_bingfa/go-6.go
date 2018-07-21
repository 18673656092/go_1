package main

type PipeData struct {
	Value int
	handler func(int) int
	next chan int
}

func handler(queue chan *PipeData){
	for data:=range queue{
		data.next <- data.handler(data.Value)
	}
}

