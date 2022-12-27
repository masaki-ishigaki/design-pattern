package builder

import (
	"log"
	"os"
)

type HTMLBuilder struct {
	fileName string
	body     string
}

func NewHTMLBuilder() Builder {
	return &HTMLBuilder{
		fileName: "untitle.html",
		body:     "",
	}
}

func (b *HTMLBuilder) MakeTitle(title string) {
	b.fileName = title + ".html"
	b.body += "<!DOCTYPE html>\n"
	b.body += "<html>\n"
	b.body += "<head><title>"
	b.body += title
	b.body += "</title></head>\n"
	b.body += "<body>\n"
	b.body += "<h1>"
	b.body += title
	b.body += "</h1>\n\n"
}

func (b *HTMLBuilder) MakeString(str string) {
	b.body += "<p>"
	b.body += str
	b.body += "</p>\n\n"
}

func (b *HTMLBuilder) MakeItems(items []string) {
	b.body += "<ul>"
	for _, item := range items {
		b.body += "<li>"
		b.body += item
		b.body += "</li>\n"
	}
	b.body += "\n"
}

func (b *HTMLBuilder) Close() {
	file, err := os.Create(b.fileName)
	if err != nil {
		log.Fatal("couldn't open file")
	}
	defer file.Close()

	b.body += "</body>"
	b.body += "</html>\n"

	byteBody := []byte(b.body)
	_, err = file.Write(byteBody)
	if err != nil {
		log.Fatal("couldn't write body to file")
	}
}

func (b *HTMLBuilder) GetHTMLResult() string {
	return b.fileName
}
