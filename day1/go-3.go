package main

import (
	"fmt"
	"reflect"
)

type Dog struct {
	Name string
	LifeExpectance int
}

func (b *Dog) Run() {
	fmt.Print("run")
}

type IRun interface {
	Run()
}

func main(){
	dog := &Dog{"dog",9}
	s := reflect.ValueOf(dog).Elem()
	typeOfT := s.Type()
	for i:=0;i<s.NumField();i++{
		f:= s.Field(i)
		fmt.Printf("%d %s %s = %v\n",i,typeOfT.Field(i).Name,f.Type(),f.Interface())
	}
}



