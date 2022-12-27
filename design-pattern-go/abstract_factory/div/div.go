package div

import (
	f "design-pattern-go/abstract_factory/factory"
)

/*****************************
 * DivLink
 ****************************/
type DivLink struct {
	f.Link
}

func NewDivLink(caption, url string) *DivLink {
	return &DivLink{
		*f.NewLink(caption, url),
	}
}

func (l *DivLink) MakeHTML() string {
	return "<div class=\"LINK\"><a href=\"" + l.Url + "\">" + l.Item.Caption + "</a></div>\n"
}

/*****************************
 * DivTray
 ****************************/
type DivTray struct {
	f.Tray
	str string
}

func NewDivTray(caption string) *DivTray {
	return &DivTray{
		Tray: *f.NewTray(caption),
		str:  "",
	}
}

func (l *DivTray) MakeHTML() string {
	l.str += "<p><b>\n"
	l.str += l.Caption
	l.str += "</b></p>\n"
	l.str += "<div class=\"TRAY\">"

	for _, item := range l.Tray.Tray {
		l.str += item.MakeHTML()
	}

	l.str += "</div>\n"

	return l.str
}

/*****************************
 * DivPage
 ****************************/
type DivPage struct {
	f.Page
	str string
}

func NewDivPage(title, author string) *DivPage {
	return &DivPage{
		Page: *f.NewPage(title, author),
		str:  "",
	}
}

func (l *DivPage) MakeHTML() string {
	l.str += "<!DOCTYPE html>\n"
	l.str += "<html><head><title>"
	l.str += l.Title
	l.str += "</title><style>\n"
	l.str += "div.TRAY { padding:0.5em; margin-left:5em; border:1px solid black; }\n"
	l.str += "div.LINK { padding:0.5em; background-color: lightgray; }\n"
	l.str += "</style></head><body>\n"
	l.str += "<h1>\n"
	l.str += l.Title
	l.str += "</h1>\n"

	for _, item := range l.Content {
		l.str += item.MakeHTML()
	}

	l.str += "<hr><address>"
	l.str += l.Author
	l.str += "</address>"
	l.str += "</body></html>\n"

	return l.str
}

/*****************************
 * ListFactory
 ****************************/
type ListFactory struct{}

func NewDivFactory() f.FactoryInterface {
	return &ListFactory{}
}

func (l *ListFactory) CreateLink(caption, url string) f.LinkInterface {
	return NewDivLink(caption, url)
}

func (l *ListFactory) CreateTray(caption string) f.TrayInterface {
	return NewDivTray(caption)
}

func (l *ListFactory) CreatePage(title, author string) f.PageInterface {
	return NewDivPage(title, author)
}
