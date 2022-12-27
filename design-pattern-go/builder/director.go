package builder

type Director struct {
	Builder *Builder
}

func NewDirector(builder *Builder) *Director {
	return &Director{
		Builder: builder,
	}
}

func (d *Director) Construct() {
	(*d.Builder).MakeTitle("Greeting")
	(*d.Builder).MakeString("一般的なあいさつ")
	(*d.Builder).MakeItems([]string{
		"How are you?",
		"Hello.",
		"Hi.",
	})
	(*d.Builder).MakeString("時間帯に応じたあいさつ")
	(*d.Builder).MakeItems([]string{
		"Good morning.",
		"Good afternoon.",
		"Good eveining.",
	})
	(*d.Builder).Close()
}
