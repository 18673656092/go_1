package main

type File struct {

}

func (f *File) Read(buf []byte)(n int,err error){
	return len(buf),err
}

func (f *File) Write(buf []byte)(n int,err error){
	return len(buf),err
}

func (f *File) Seek(off int64,whence int)(pos int64,err error){
	return off,err
}

type IRead interface {
	Read(buf []byte)(n int,err error)
}

type TWrite interface {
	Write(buf []byte)(n int,err error)
}
