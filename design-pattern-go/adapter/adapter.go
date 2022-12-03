package adapter

import "fmt"

// Printはなし

/*****************************
 * Banner
 ****************************/
type Banner struct {
	str string
}

func NewBanner(str string) *Banner {
	return &Banner{
		str,
	}
}

func (b *Banner) ShowWithParen() {
	fmt.Printf("(%s)\n", b.str)
}

func (b *Banner) ShowWithAster() {
	fmt.Printf("*%s*\n", b.str)
}

/*****************************
 * PrintBanner
 ****************************/
type PrintBanner struct {
	banner *Banner
}

func NewPrintBanner(banner *Banner) *PrintBanner {
	return &PrintBanner{
		banner,
	}
}

func (p *PrintBanner) PrintWeak() {
	p.banner.ShowWithParen()
}

func (p *PrintBanner) PrintStrong() {
	p.banner.ShowWithAster()
}
