package factory

import (
	"fmt"
	"log"
	"os"
)

/*****************************
 * Item
 ****************************/
type ItemInterface interface {
	MakeHTML() string
}

type Item struct {
	Caption string
}

/*****************************
 * Link
 ****************************/
type LinkInterface interface {
	ItemInterface
}

type Link struct {
	Item
	Url string
}

func NewLink(caption, url string) *Link {
	return &Link{
		Item: Item{
			caption,
		},
		Url: url,
	}
}

/*****************************
 * Tray
 ****************************/
type TrayInterface interface {
	ItemInterface
	Add(item ItemInterface)
}

type Tray struct {
	Item
	Tray []ItemInterface
}

func NewTray(caption string) *Tray {
	return &Tray{
		Item: Item{
			caption,
		},
		Tray: []ItemInterface{},
	}
}

func (t *Tray) Add(item ItemInterface) {
	t.Tray = append(t.Tray, item)
}

/*****************************
 * Page
 ****************************/
type PageInterface interface {
	MakeHTML() string
	Add(item ItemInterface)
	Output(filename string, page PageInterface)
}

type Page struct {
	Title   string
	Author  string
	Content []ItemInterface
}

func NewPage(title, author string) *Page {
	return &Page{
		Title:   title,
		Author:  author,
		Content: []ItemInterface{},
	}
}

func (p *Page) Add(item ItemInterface) {
	p.Content = append(p.Content, item)
}

func (p *Page) Output(filename string, page PageInterface) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("couldn't open file")
	}
	defer file.Close()

	byteBody := []byte(page.MakeHTML())
	_, err = file.Write(byteBody)
	if err != nil {
		log.Fatal("couldn't write body to file")
	}

	fmt.Println(filename + "を作成しました。")

}

/*****************************
 * Factory
 ****************************/
type FactoryInterface interface {
	CreateLink(caption, url string) LinkInterface
	CreateTray(caption string) TrayInterface
	CreatePage(title, author string) PageInterface
}
