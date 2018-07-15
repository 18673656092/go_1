package main

var (
	v1 int
	v2 string
	v3 [10]int
	v4 []int
	v5 struct{
		f int
	}
	v6 *int
	v7 map[string]int
	v8 func(i int)(int)
)

func main(){
	v1 = 1
	v9 := 3
	v1,v9 = v9,v1
	println(v1,v9)

}


