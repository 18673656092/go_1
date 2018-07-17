package main

import "fmt"

func f1(){
	var messages chan string = make(chan string)
	go func(message string) {messages <- message}("PING")
	fmt.Print(<-messages)
}

var ch chan int = make(chan int)

func f2(){
	func() { ch<-0}()
}

func main(){
	go f2()
	<-ch
}

