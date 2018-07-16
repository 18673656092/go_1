package main

import "fmt"

type Base struct {
	Name string
}

func (base *Base) Foo(){

}

func (base *Base) Bar(){

}

type Foo struct {
	Base
}

func (foo *Foo) Bar()  {
	foo.Base.Bar()
}

func main(){
	x := Rect{1,2,3,4,0}
	fmt.Print(x)
}