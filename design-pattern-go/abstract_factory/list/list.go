package list

import (
	f "design-pattern-go/abstract_factory/factory"
)

/*****************************
 * ListLink
 ****************************/
type ListLink struct {
	f.Link
}

func NewListLink(caption, url string) *ListLink {
	return &ListLink{
		*f.NewLink(caption, url),
	}
}

func (l *ListLink) MakeHTML() string {
	return "  <li><a href=\"" + l.Url + "\">" + l.Item.Caption + "</a></li>\n"
}

/*****************************
 * ListTray
 ****************************/
type ListTray struct {
	f.Tray
	str string
}

func NewListTray(caption string) *ListTray {
	return &ListTray{
		Tray: *f.NewTray(caption),
		str:  "",
	}
}

func (l *ListTray) MakeHTML() string {
	l.str += "<li>\n"
	l.str += l.Caption
	l.str += "\n<ul>\n"

	for _, item := range l.Tray.Tray {
		l.str += item.MakeHTML()
	}

	l.str += "</ul>\n"
	l.str += "</li>\n"

	return l.str
}

/*****************************
 * ListPage
 ****************************/
type ListPage struct {
	f.Page
	str string
}

func NewListPage(title, author string) *ListPage {
	return &ListPage{
		Page: *f.NewPage(title, author),
		str:  "",
	}
}

func (l *ListPage) MakeHTML() string {
	l.str += "<!DOCTYPE html>\n"
	l.str += "<html><head><title>"
	l.str += l.Title
	l.str += "</title></head>\n"
	l.str += "<body>\n"
	l.str += "<h1>\n"
	l.str += l.Title
	l.str += "</h1>\n"
	l.str += "<ul>\n"

	for _, item := range l.Content {
		l.str += item.MakeHTML()
	}

	l.str += "</ul>\n"
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

func NewListFactory() f.FactoryInterface {
	return &ListFactory{}
}

func (l *ListFactory) CreateLink(caption, url string) f.LinkInterface {
	return NewListLink(caption, url)
}

func (l *ListFactory) CreateTray(caption string) f.TrayInterface {
	return NewListTray(caption)
}

func (l *ListFactory) CreatePage(title, author string) f.PageInterface {
	return NewListPage(title, author)
}
