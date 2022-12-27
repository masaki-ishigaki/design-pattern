package display

import (
	di "design-pattern-go/bridge/display_impl"
)

type Display struct {
	impl di.DisplyImpl
}

func NewDisplay(impl di.DisplyImpl) *Display {
	return &Display{
		impl,
	}
}

func (d *Display) Open() {
	d.impl.RawOpen()
}

func (d *Display) Print() {
	d.impl.RawPrint()
}

func (d *Display) Close() {
	d.impl.RawClose()
}

func (d *Display) Display() {
	d.Open()
	d.Print()
	d.Close()
}

type CountDisplay struct {
	Display
}

func NewCountDisplay(impl di.DisplyImpl) *CountDisplay {
	return &CountDisplay{
		*NewDisplay(impl),
	}
}

func (c *CountDisplay) MultiDisplay(times int) {
	c.Open()
	for i := 0; i < times; i++ {
		c.Print()
	}
	c.Close()
}
