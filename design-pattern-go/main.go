package main

import (
	"design-pattern-go/adapter"
	fm "design-pattern-go/factory_method"
	iter "design-pattern-go/iterator"
	pr "design-pattern-go/prototype"
	"design-pattern-go/singleton"
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

func trySingleton() {
	fmt.Println("Start")
	obj1 := singleton.GetInstance()
	obj2 := singleton.GetInstance()
	if obj1 == obj2 {
		fmt.Println("obj1とobj2は同じインスタンスです。")
	} else {
		fmt.Println("obj1とobj2は同じインスタンスではありません。")
	}
	fmt.Println("End")
}

func tryPrototype() {
	manager := pr.NewManager()
	upen := pr.NewUnderlinePen("_")
	mbox := pr.NewMessageBox("*")
	sbox := pr.NewMessageBox("/")

	manager.Register("strong message", upen)
	manager.Register("warning box", mbox)
	manager.Register("slash box", sbox)

	p1 := manager.Create("strong message")
	p1.Use("Hello, world.")

	p2 := manager.Create("warning box")
	p2.Use("Hello, world.")

	p3 := manager.Create("slash box")
	p3.Use("Hello, world.")
}

func main() {
	tryPrototype()
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/http/httputil"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
// 	dump, err := httputil.DumpRequest(r, true)
// 	if err != nil {
// 		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Println(string(dump))
// 	fmt.Fprintf(w, "<html><body>hello<body><html>\n")
// }

// func main() {
// 	var httpServer http.Server
// 	http.HandleFunc("/", handler)
// 	log.Println("start http listening :1888")
// 	httpServer.Addr = ":18888"
// 	log.Println(httpServer.ListenAndServe())
// }
