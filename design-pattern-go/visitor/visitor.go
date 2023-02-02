package visitor

import "fmt"

type Visitor interface {
	VisitFile(file File)
	VisitDirectory(directory Directory)
}

type ListVisitor struct {
	currentDir string
}

func NewListVisitor() *ListVisitor {
	return &ListVisitor{
		currentDir: "",
	}
}

func (l *ListVisitor) VisitFile(file File) {
	fmt.Printf("%s/%v\n", l.currentDir, file.Entry.concreteEntry)
}

func (l *ListVisitor) VisitDirectory(dir Directory) {
	fmt.Printf("%s/%v\n", l.currentDir, dir.Entry.concreteEntry)
	savedir := l.currentDir
	l.currentDir = l.currentDir + "/" + dir.GetName()
	for _, entry := range dir.directory {
		entry.element.Accept(l)
	}
	l.currentDir = savedir
}
