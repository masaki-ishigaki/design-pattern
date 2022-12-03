package template_method

import "fmt"

type Printer interface {
	open()
	print()
	close()
}

/*****************************
 * AbstractDisplay
 ****************************/
type AbstractDisplay struct {
	printer Printer
}

func (a *AbstractDisplay) Display() {
	a.printer.open()
	for i := 0; i < 5; i++ {
		a.printer.print()
	}
	a.printer.close()
}

/*****************************
 * CharDisplay
 ****************************/
type CharDisplay struct {
	*AbstractDisplay
	ch rune
}

func NewCharDispaly(ch rune) *CharDisplay {
	charDisplay := &CharDisplay{
		AbstractDisplay: &AbstractDisplay{},
		ch:              ch,
	}
	charDisplay.printer = charDisplay
	return charDisplay
}

func (c *CharDisplay) open() {
	fmt.Print("<<")
}

func (c *CharDisplay) print() {
	fmt.Print(string(c.ch))
}

func (c *CharDisplay) close() {
	fmt.Print(">>")
}

/*****************************
 * StringDisplay
 ****************************/
type StringDisplay struct {
	*AbstractDisplay
	str   string
	width int
}

func NewStringDispaly(str string) *StringDisplay {
	stringDisplay := &StringDisplay{
		AbstractDisplay: &AbstractDisplay{},
		str:             str,
		width:           len(str),
	}
	stringDisplay.printer = stringDisplay
	return stringDisplay
}

func (s *StringDisplay) open() {
	s.printLine()
}

func (s *StringDisplay) print() {
	fmt.Printf("%s%s%s\n", "|", s.str, "|")
}

func (s *StringDisplay) close() {
	s.printLine()
}

func (s *StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < s.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
