package factory_method

import "fmt"

/*****************************
 * Product
 ****************************/
type Product interface {
	Use()
	GetOwner() string
}

/*****************************
 * Factory
 ****************************/
type IFactory interface {
	createProduct(owner string) Product
	registerProduct(product Product)
}

type Factory struct {
	factory IFactory
}

func (f *Factory) Create(owner string) Product {
	p := f.factory.createProduct(owner)
	f.factory.registerProduct(p)
	return p
}

/*****************************
 * IDCard
 ****************************/
type iDCard struct {
	owner string
}

func newIDCard(owner string) *iDCard {
	fmt.Printf("%sのカードを作ります。", owner)
	fmt.Println("")
	return &iDCard{
		owner,
	}
}

func (i *iDCard) Use() {
	fmt.Println(i, "を使います。")
}

func (i *iDCard) String() string {
	return fmt.Sprintf("[IDCard:%s]", i.owner)
}

func (i *iDCard) GetOwner() string {
	return i.owner
}

/*****************************
 * IDCard Factory
 ****************************/
type IDCardFactory struct {
	*Factory
}

func NewIDCardFactory() *IDCardFactory {
	idCardFactory := &IDCardFactory{
		Factory: &Factory{},
	}
	idCardFactory.factory = idCardFactory
	return idCardFactory
}

func (i *IDCardFactory) createProduct(owner string) Product {
	return newIDCard(owner)
}

func (i *IDCardFactory) registerProduct(product Product) {
	fmt.Println(product, "を登録しました。")
}
