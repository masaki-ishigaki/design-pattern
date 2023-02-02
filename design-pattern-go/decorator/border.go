package decorator

type Border struct {
	display
	target IDisplay
}

func NewBorder(disp IDisplay) *Border {
	return &Border{
		target: disp,
	}
}
