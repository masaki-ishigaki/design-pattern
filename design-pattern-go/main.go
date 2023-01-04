package main

import (
	af "design-pattern-go/abstract_factory"
	"design-pattern-go/adapter"
	disp "design-pattern-go/bridge/display"
	dispi "design-pattern-go/bridge/display_impl"
	"design-pattern-go/builder"
	fm "design-pattern-go/factory_method"
	iter "design-pattern-go/iterator"
	pr "design-pattern-go/prototype"
	"design-pattern-go/singleton"
	"design-pattern-go/strategy"
	tm "design-pattern-go/template_method"
	"fmt"
	"os"
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

func tryBuilder(builderType string) {
	if builderType == "text" {
		textBuilder := builder.NewTextBuilder()
		director := builder.NewDirector(&textBuilder)
		director.Construct()

		castedTextBuilder := textBuilder.(*builder.TextBuilder)
		fmt.Println(castedTextBuilder.GetTextResult())
	} else if builderType == "html" {
		htmlBuilder := builder.NewHTMLBuilder()
		director := builder.NewDirector(&htmlBuilder)
		director.Construct()

		castedHTMLBuilder := htmlBuilder.(*builder.HTMLBuilder)
		fmt.Println(castedHTMLBuilder.GetHTMLResult())
	} else {
		fmt.Print("Inappropriate builder type is input!!")
	}
}

func tryAbstractFactory(fileName, factoryName string) {
	factory, _ := af.GetFactory(factoryName)

	// Blog
	blog1 := factory.CreateLink("Blog 1", "https://example.com/blog1")
	blog2 := factory.CreateLink("Blog 2", "https://example.com/blog2")
	blog3 := factory.CreateLink("Blog 3", "https://example.com/blog3")

	blogTray := factory.CreateTray("Blog Site")
	blogTray.Add(blog1)
	blogTray.Add(blog2)
	blogTray.Add(blog3)

	// News
	news1 := factory.CreateLink("News 1", "https://example.com/news1")
	news2 := factory.CreateLink("News 2", "https://example.com/news2")
	news3 := factory.CreateTray("News 3")
	news3.Add(factory.CreateLink("News 3 (US)", "https://example.com/news3us"))
	news3.Add(factory.CreateLink("News 3 (JP)", "https://example.com/news3jp"))

	newsTray := factory.CreateTray("News Site")
	newsTray.Add(news1)
	newsTray.Add(news2)
	newsTray.Add(news3)

	// Page
	page := factory.CreatePage("Blog and News", "Hiroshi Yuki")
	page.Add(blogTray)
	page.Add(newsTray)

	page.Output(fileName, page)
}

func tryBridge() {
	d1 := disp.NewDisplay(dispi.NewStringDisplayImpl("Hello, Japan."))
	d2 := disp.NewCountDisplay(dispi.NewStringDisplayImpl("Hello, World."))
	d3 := disp.NewCountDisplay(dispi.NewStringDisplayImpl("Hello, Universe."))

	d1.Display()
	d2.Display.Display()
	d3.Display.Display()
	d3.MultiDisplay(5)
}

func tryStrategy() {
	player1 := strategy.NewPlayer("Taro", strategy.NewWinningStrategy())
	player2 := strategy.NewPlayer("Hana", strategy.NewProbStrategy())

	for i := 0; i < 10000; i++ {
		nextHand1 := player1.NextHand()
		nextHand2 := player2.NextHand()

		if nextHand1.IsStrongerThan(nextHand2) {
			fmt.Println("Winner:", player1)
			player1.Win()
			player2.Lose()
		} else if nextHand2.IsStrongerThan(nextHand1) {
			fmt.Println("Winner:", player2)
			player1.Lose()
			player2.Win()
		} else {
			fmt.Println("Even...")
			player1.Even()
			player2.Even()
		}
	}

	fmt.Println("Total result:")
	fmt.Println(player1)
	fmt.Println(player2)
}

func main() {
	pattern := os.Args[1]

	switch pattern {
	case "iterator":
		tryIterator()
	case "adapter":
		tryAdapter()
	case "template_method":
		tryTemplateMethod()
	case "factory_method":
		tryFactoryMethod()
	case "singleton":
		trySingleton()
	case "prototype":
		tryPrototype()
	case "builder":
		tryBuilder(os.Args[2])
	case "abstract_factory":
		tryAbstractFactory(os.Args[2], os.Args[3])
	case "bridge":
		tryBridge()
	case "strategy":
		tryStrategy()
	default:
		fmt.Println("chose proper pattern!!")
	}
}
