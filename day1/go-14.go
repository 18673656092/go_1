package main

import (
	"errors"
	"fmt"
)

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Error")
		return 0,err
	}
	return a + b, nil
}
func myfunc(args ...interface{}){
	for _, arg:= range args{
		switch arg.(type) {
		case int: fmt.Println(arg,"is an int value")
		case string: fmt.Println(arg,"is a string value")
		default: fmt.Println(arg,"is an unknown type")
		}
	}
}
func main() {
	myfunc(1,2,(int16(3)),4,5,6,"2sd")
}
