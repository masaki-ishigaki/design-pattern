package facade

import (
	"fmt"
	"os"
)

type HtmlWriter struct {
	writer *os.File
}

func NewHtmlWriter(writer *os.File) *HtmlWriter {
	return &HtmlWriter{
		writer,
	}
}

func (h *HtmlWriter) Title(title string) {
	h.writer.WriteString("<!DOCTYPE html")
	h.writer.WriteString("<html>")
	h.writer.WriteString("<head>")
	h.writer.WriteString(fmt.Sprintf("<title>%s</title>", title))
	h.writer.WriteString("</head>")
	h.writer.WriteString("<body>")
	h.writer.WriteString("\n")
	h.writer.WriteString(fmt.Sprintf("<h1>%s</h1>", title))
	h.writer.WriteString("\n")
}

func (h *HtmlWriter) Paragraph(msg string) {
	h.writer.WriteString(fmt.Sprintf("<p>%s</p>\n", msg))
}

func (h *HtmlWriter) Link(href, caption string) {
	h.Paragraph(fmt.Sprintf("<a href=\"%s\">%s</a>", href, caption))
}

func (h *HtmlWriter) MailTo(mailaddr, username string) {
	h.Link(fmt.Sprintf("mailto: %s", mailaddr), username)
}

func (h *HtmlWriter) Close() {
	h.writer.WriteString("</body>")
	h.writer.WriteString("</html>")
	h.writer.WriteString("\n")
	h.writer.Close()
}
