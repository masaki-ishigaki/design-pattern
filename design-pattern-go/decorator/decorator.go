package decorator

import (
	"fmt"
)

type IDisplay interface {
	GetColumns() int
	GetRows() int
	GetRowText(row int) string
}

type display struct {
	concreteDisp IDisplay
}

func (d *display) Show() {
	for i := 0; i < d.concreteDisp.GetRows(); i++ {
		fmt.Println(d.concreteDisp.GetRowText(i))
	}
}
