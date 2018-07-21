package main

import "fmt"

type File struct {
}

func (f *File) Read(buf []byte) (n int, err error) {
	return len(buf), err
}

func (f *File) Write(buf []byte) (n int, err error) {
	return len(buf), err
}

func (f *File) Seek(off int64, whence int) (pos int64, err error) {
	return off, err
}

type IRead interface {
	Read(buf []byte) (n int, err error)
}

type TWrite interface {
	Write(buf []byte) (n int, err error)
}

type IWriter interface {
	Write(buf []byte) (n int, err error)
}

func main(){
	var file1 IRead = new(File)
	t := []byte{1,2,3,4,5,6,7}
	fmt.Println(file1.Read(t))
	if file,ok := file1.(*File);ok{
		fmt.Println("OK",file)
	}
}