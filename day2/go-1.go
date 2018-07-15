package main

import "fmt"

type Interger int

func (o Interger) Less(p Interger) bool  {
	return o < p
}
func (o *Interger) Add(p Interger) Interger {
	*o += p
	return *o
}
func (o Interger) AddOne() Interger{
	o += 1
	return o
}
func main(){
	var a Interger = 1
	if a.Less(2){
		fmt.Println(a,"Less 2")
	}
	fmt.Println(a.Add(3))
	fmt.Println(a.AddOne())
	fmt.Println(a)
}