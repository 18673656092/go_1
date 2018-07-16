package main

import "fmt"

var value1 int32
var str = "hello world"

func a()  {
	for i:=0;i < len(str);i++{
		ch := str[i]
		fmt.Print(ch + ' ')
		fmt.Print(string(ch)+" ")
	}
}

func main()  {
	value2 := 64
	value1 = int32(value2)
	//fmt.Print(str[0])
	//fmt.Print(len(str))
	a()
}
