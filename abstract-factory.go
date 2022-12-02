package main

/* abstract factory */

import "fmt"

// define an abstract factory that makes abstract products
type iAbstractFactory interface {
	makeProduct() iAbstractProduct
}

// free-standing abstract factory constructor
func newAbstractFactory(factoryName string) (iAbstractFactory, error) {
	var newFactory = new(concreteFactory)
	newFactory.name = factoryName
	return newFactory, nil
}

// define an abstract product (A) with some methods
type iAbstractProduct interface {
	setPropertyA(attribute string)
	getPropertyA() string
	setSizeOfA(size int)
	getSizeOfA() int
}

// abstract product (A) method constructor
func (*concreteFactory) makeProduct() iAbstractProduct {
	var produceA = new(concreteProduct)
	return produceA
}

// define a concrete factory type
type concreteFactory struct {
	name string
}

// define a concrete product (A) type
type concreteProduct struct {
	attribute string
	size      int
}

// concrete productA implementations
// ---------------------------------

func (item *concreteProduct) setPropertyA(attribute string) {
	item.attribute = attribute
}

func (item *concreteProduct) getPropertyA() string {
	return item.attribute
}

func (item *concreteProduct) setSizeOfA(size int) {
	item.size = size
}

func (s *concreteProduct) getSizeOfA() int {
	return s.size
}

func main() {
	// create a new abstract factory
	myFactory, _ := newAbstractFactory("myAwesomeBrand")

	// create productA from abstract factory
	myProduct := myFactory.makeProduct()
	makeAbstractProduct(myProduct)

}

func makeAbstractProduct(prodA iAbstractProduct) {

	prodA.setPropertyA("Set product A attribute here.")
	prodA.setSizeOfA(69)
	fmt.Printf("Product A Property: %s\n", prodA.getPropertyA())
	fmt.Printf("Product A Size     : %d\n", prodA.getSizeOfA())
}
