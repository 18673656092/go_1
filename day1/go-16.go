package main

import (
	"os"
	"fmt"
)

type PathError struct {
	Op string
	Path string
	Err error
}

func (e *PathError) Error() string{
	return e.Op + " " + e.Path + " " + e.Err.Error()
}

func main(){
	fi,err := os.Stat("a.txt")
	if err != nil{
		if e,ok := err.(*os.PathError); ok && e.Err != nil{
			fmt.Println(fi)
			fmt.Println(e.Err)
		}
	}
	defer fmt.Println("1")
	defer fmt.Println("2")
}
