package main

import (
	"design-pattern-go/adapter"
	fm "design-pattern-go/factory_method"
	iter "design-pattern-go/iterator"
	tm "design-pattern-go/template_method"
	"fmt"
)

func tryIterator() {
	bookShelf := iter.NewBookShelf(4)
	bookShelf.AppendBook(iter.NewBook("Around the World in 80 days"))
	bookShelf.AppendBook(iter.NewBook("Bible"))
	bookShelf.AppendBook(iter.NewBook("Cindrella"))
	bookShelf.AppendBook(iter.NewBook("Daddy-Long-Legs"))

	it := bookShelf.Iterator()
	for it.HasNext() {
		book := it.Next()
		fmt.Println(book)
	}
}

func tryAdapter() {
	p := adapter.NewPrintBanner(adapter.NewBanner("Hello"))
	p.PrintWeak()
	p.PrintStrong()
}

func tryTemplateMethod() {
	d1 := tm.NewCharDispaly('H')
	d2 := tm.NewStringDispaly("Hello, world.")
	d1.Display()
	fmt.Println("")
	d2.Display()
}

func tryFactoryMethod() {
	factory := fm.NewIDCardFactory()
	card1 := factory.Create("Hiroshi Yuki")
	card2 := factory.Create("Tomura")
	card3 := factory.Create("Hanako Sato")

	card1.Use()
	card2.Use()
	card3.Use()
}

func main() {
	tryFactoryMethod()
}
