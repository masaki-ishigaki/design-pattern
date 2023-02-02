package visitor

type Directory struct {
	Entry
	name      string
	directory []Entry
}

func NewDirectory(name string) *Directory {
	dir := &Directory{
		name:      name,
		directory: []Entry{},
	}
	dir.Entry.element = dir
	dir.Entry.concreteEntry = dir

	return dir
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) GetSize() int {
	size := 0
	for _, entry := range d.directory {
		size += entry.concreteEntry.GetSize()
	}
	return size
}

func (d *Directory) Accept(v Visitor) {
	v.VisitDirectory(*d)
}

func (d *Directory) Add(entry Entry) {
	d.directory = append(d.directory, entry)
}
