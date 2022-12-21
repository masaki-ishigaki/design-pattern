package prototype

import "fmt"

type MessageBox struct {
	decochar string
}

func NewMessageBox(decochar string) *MessageBox {
	return &MessageBox{
		decochar,
	}
}

func (m *MessageBox) Use(s string) {
	decolen := 1 + len(s) + 1

	for i := 0; i < decolen; i++ {
		fmt.Print(m.decochar)
	}
	fmt.Println("")
	fmt.Printf("%s%s%s", m.decochar, s, m.decochar)
	fmt.Println("")
	for i := 0; i < decolen; i++ {
		fmt.Print(m.decochar)
	}
	fmt.Println("")
}

func (m *MessageBox) CreateCopy() Product {
	return NewMessageBox(m.decochar)
}
