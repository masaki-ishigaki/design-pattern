package builder

type TextBuilder struct {
	body string
}

func NewTextBuilder() Builder {
	return &TextBuilder{
		body: "",
	}
}

func (b *TextBuilder) MakeTitle(title string) {
	b.body += "======================================\n"
	b.body += "『"
	b.body += title
	b.body += "』\n\n"
}

func (b *TextBuilder) MakeString(str string) {
	b.body += "■"
	b.body += str
	b.body += "\n\n"
}

func (b *TextBuilder) MakeItems(items []string) {
	for _, item := range items {
		b.body += "  ・"
		b.body += item
		b.body += "\n"
	}
	b.body += "\n"
}

func (b *TextBuilder) Close() {
	b.body += "======================================\n"
}

func (b *TextBuilder) GetTextResult() string {
	return b.body
}
