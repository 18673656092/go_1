package main

import "fmt"

func modfiy(array [5]int){
	array[0] = 10
	fmt.Println("In modify(),array values:",array)
}

func main()  {
	array := [5]int{1,2,3,4,5}
	modfiy(array)
	fmt.Println("In main(),array values:",array)
}
