package composite

import "fmt"

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{
		name,
		size,
	}
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) PrintList(prefix string) {
	fmt.Printf("%s/%s(%d)\n", prefix, f.GetName(), f.GetSize())
}
