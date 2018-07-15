package main

func main(){
	const (
		c0 = iota
		c1 = iota
		c2 = iota
	)
	println(c0,c1,c2)
	println(c0,c1,c2,c0)
}
