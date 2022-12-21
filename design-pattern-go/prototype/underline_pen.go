package prototype

import "fmt"

type UnderlinePen struct {
	ulchar string
}

func NewUnderlinePen(ulchar string) *UnderlinePen {
	return &UnderlinePen{
		ulchar,
	}
}

func (u *UnderlinePen) Use(s string) {
	ulen := len(s)

	fmt.Println(s)
	for i := 0; i < ulen; i++ {
		fmt.Print(u.ulchar)
	}
	fmt.Println("")
}

func (u *UnderlinePen) CreateCopy() Product {
	return NewUnderlinePen(u.ulchar)
}
