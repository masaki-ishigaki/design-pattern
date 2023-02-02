package visitor

import "fmt"

type Element interface {
	Accept(v Visitor)
}

type IEntry interface {
	GetName() string
	GetSize() int
}

type Entry struct {
	element       Element
	concreteEntry IEntry
}

func NewEntry() *Entry {
	return &Entry{}
}

func (e *Entry) String() string {
	return fmt.Sprintf("%s (%d)", e.concreteEntry.GetName(), e.concreteEntry.GetSize())
}
