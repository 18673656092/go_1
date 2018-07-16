package main

import "fmt"

type Int int

func (a Int) Less(b Int) bool {
	return a < b
}

func (a *Int) Add(b Int) {
	*a += b
}

type LessAndAdd interface {
	Less(b Int) bool
	Add(b Int)
}

type Lesser interface {
	Less(b Int) bool
} 

func main(){
	var a Int = 1
	var b LessAndAdd = &a
	var c Lesser = a
	var d Lesser = &a
	b.Add(3)
	fmt.Print(b.Less(3),c.Less(3),d.Less(3),b,c,d)
}

