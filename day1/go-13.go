package main

import "fmt"

func example(x int)(int){
	if x==0 {
		return x
	}else if x==9 {
		sum := 0
		for i := 1;i < x ;i++  {
			if i==6{
				break
			}else {
				sum += i
			}
		}
		return sum
	}else {
		switch x{
		case 1:
			return 1
		case 2:
			return 2
		default:
			return 0
		}
	}
}

func main()  {
	fmt.Print(example(9))
}


