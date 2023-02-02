package composite

type Entry interface {
	GetName() string
	GetSize() int
	PrintList(prefix string)
}
