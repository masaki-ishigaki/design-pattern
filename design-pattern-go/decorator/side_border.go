package decorator

type SideBorder struct {
	borderCh string
	Border
}

func NewSideBorder(disp IDisplay, ch string) *SideBorder {
	sideBorder := &SideBorder{
		borderCh: ch,
		Border:   *NewBorder(disp),
	}
	// interfaceの本実装
	sideBorder.Border.display.concreteDisp = sideBorder

	return sideBorder
}

func (s *SideBorder) GetColumns() int {
	return 1 + s.Border.target.GetColumns() + 1
}

func (s *SideBorder) GetRows() int {
	return s.Border.target.GetRows()
}

func (s *SideBorder) GetRowText(row int) string {
	return s.borderCh + s.Border.target.GetRowText(row) + s.borderCh
}
