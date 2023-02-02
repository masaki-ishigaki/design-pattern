package decorator

type FullBorder struct {
	Border
}

func NewFullBorder(disp IDisplay) *FullBorder {
	fullBorder := &FullBorder{
		Border: *NewBorder(disp),
	}
	// interfaceの本実装
	fullBorder.Border.display.concreteDisp = fullBorder

	return fullBorder
}

func (f *FullBorder) GetColumns() int {
	return 1 + f.Border.target.GetColumns() + 1
}

func (f *FullBorder) GetRows() int {
	return 1 + f.Border.target.GetRows() + 1
}

func (f *FullBorder) GetRowText(row int) string {
	if row == 0 || row == (f.Border.target.GetRows()+1) {
		return "+" + f.makeLine("-", f.Border.target.GetColumns()) + "+"
	} else {
		return "|" + f.Border.target.GetRowText(row-1) + "|"
	}
}

func (f *FullBorder) makeLine(ch string, count int) string {
	str := ""
	for i := 0; i < count; i++ {
		str += ch
	}

	return str
}
