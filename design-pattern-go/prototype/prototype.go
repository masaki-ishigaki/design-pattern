package prototype

type Product interface {
	Use(s string)
	CreateCopy() Product
}
