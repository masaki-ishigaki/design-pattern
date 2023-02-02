package visitor

type File struct {
	Entry
	name string
	size int
}

func NewFile(name string, size int) *File {
	file := &File{
		name: name,
		size: size,
	}
	file.Entry.concreteEntry = file
	file.Entry.element = file

	return file
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) Accept(v Visitor) {
	v.VisitFile(*f)
}
