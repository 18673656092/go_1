package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title string
	Authors []string
	Price float64
	List List
}

type List struct {
	Name string
	ID int
}

func main(){
	var array = []string{"test","one","two"}
	list := List{"test",1}
	gobook := Book{"GO",array,10.0,list}
	b,err := json.Marshal(gobook)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(b))
	var book Book
	e := json.Unmarshal(b,&book)
	if e != nil{
		fmt.Println(e)
	}
	fmt.Println(book)
}
