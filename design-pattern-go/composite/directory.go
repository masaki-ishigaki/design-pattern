package composite

import "fmt"

type Directory struct {
	name      string
	directory []Entry
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name:      name,
		directory: []Entry{},
	}
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) GetSize() int {
	size := 0
	for _, entry := range d.directory {
		size += entry.GetSize()
	}
	return size
}

func (d *Directory) PrintList(prefix string) {
	fmt.Printf("%s/%s(%d)\n", prefix, d.GetName(), d.GetSize())
	for _, entry := range d.directory {
		entry.PrintList(prefix + "/" + d.name)
	}
}

func (d *Directory) Add(entry Entry) {
	d.directory = append(d.directory, entry)
}
