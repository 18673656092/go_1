package main

import "fmt"

func Tel()([] int){
	mySlice := make([]int,5,10)
	return mySlice
}

func main(){
	myArray := Tel()
	fmt.Println(cap(myArray))
	fmt.Println(len(myArray))
}