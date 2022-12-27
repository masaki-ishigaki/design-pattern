package display_impl

import "fmt"

type DisplyImpl interface {
	RawOpen()
	RawPrint()
	RawClose()
}

type StringDisplayImpl struct {
	str   string
	width int
}

func NewStringDisplayImpl(str string) DisplyImpl {
	return &StringDisplayImpl{
		str:   str,
		width: len(str),
	}
}

func (s *StringDisplayImpl) RawOpen() {
	s.printLine()
}

func (s *StringDisplayImpl) RawPrint() {
	output := "|" + s.str + "|"
	fmt.Println(output)
}

func (s *StringDisplayImpl) RawClose() {
	s.printLine()
}

func (s *StringDisplayImpl) printLine() {
	fmt.Print("+")
	for i := 0; i < s.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
