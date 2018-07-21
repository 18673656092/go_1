package main

import "fmt"

type Rect struct {
	X,Y float64
	Width,Height float64
	S float64
}

type Float float64

func (r *Rect) Area() float64{
	r.S = r.Width * r.Height
	return r.S
}

func main(){
	r := Rect{1,2,3,4,0}
	x := r.Area()
	x = r.Area()
	fmt.Println(x)
	fmt.Println(r.S)
	//a := x
	//c := &x
	////b := *x
	//fmt.Print(a,c)
}

