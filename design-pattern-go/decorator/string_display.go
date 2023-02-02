package decorator

type StringDisplay struct {
	str string
	display
}

func NewStringDisplay(str string) *StringDisplay {
	strDisplay := &StringDisplay{
		str:     str,
		display: display{},
	}
	// interfaceの本実装
	strDisplay.display.concreteDisp = strDisplay

	return strDisplay
}

func (s *StringDisplay) GetColumns() int {
	return len(s.str)
}

func (s *StringDisplay) GetRows() int {
	return 1
}

func (s *StringDisplay) GetRowText(row int) string {
	if row != 0 {
		panic("Index Error!")
	}
	return s.str
}
