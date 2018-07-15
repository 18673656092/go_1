package main

import "fmt"

type Bird struct {
	i string
}

func (b *Bird) Fly(){
	b.i = "hello"
	fmt.Print(b.i)
}

type IFly interface {
	Fly()
}

func main() {
	var fly IFly = new(Bird)
	fly.Fly()
}