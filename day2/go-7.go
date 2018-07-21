package main

import (
	"go_1/one"
	"go_1/two"
	"fmt"
)

func main() {
	var file1 one.ReadWriter = new(File)
	var file2 two.IStream = new(File)
	var file3 one.ReadWriter = file2
	var file4 two.IStream = file1
	fmt.Print(file3, file4)
}
