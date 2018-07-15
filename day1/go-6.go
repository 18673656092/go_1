package main

import "fmt"

func GetName()(firstName string,lastName string,nickName string)  {
	return "May","Chan","Chibi"
}

func main()  {
	_,_,nickName := GetName()
	fmt.Print(nickName)
}