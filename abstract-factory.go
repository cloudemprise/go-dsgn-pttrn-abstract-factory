package main

/* abstract factory */

import "fmt"

// define an abstract factory that makes abstract products
type iAbstractFactory interface {
	makeProduct() iAbstractProduct
}

// a free-standing abstract factory constructor
func newAbstractFactory(factoryName string) (iAbstractFactory, error) {
	var newFactory = new(concreteFactory)
	newFactory.name = factoryName
	return newFactory, nil
}

// define an abstract product with some behaviours
type iAbstractProduct interface {
	setAttribute(attribute string)
	getAttribute() string
	setSize(size int)
	getSize() int
}

// an abstract product constructor
func (*concreteFactory) makeProduct() iAbstractProduct {
	var produce = new(concreteProduct)
	return produce
}

// define a concrete factory type
type concreteFactory struct {
	name string
}

// define a concrete product type
type concreteProduct struct {
	attribute string
	size      int
}

// concrete product implementations
// ---------------------------------

func (item *concreteProduct) setAttribute(attribute string) {
	item.attribute = attribute
}

func (item *concreteProduct) getAttribute() string {
	return item.attribute
}

func (item *concreteProduct) setSize(size int) {
	item.size = size
}

func (item *concreteProduct) getSize() int {
	return item.size
}

func main() {
	// create a new abstract factory
	myFactory, _ := newAbstractFactory("myAwesomeBrand")

	// create product from abstract factory
	myProduct := myFactory.makeProduct()
	makeAbstractProduct(myProduct)

}

// Test Case
func makeAbstractProduct(prod iAbstractProduct) {

	prod.setAttribute("Special Attribute")
	prod.setSize(69)
	fmt.Printf("Product Attribute : %s\n", prod.getAttribute())
	fmt.Printf("Product Size     : %d\n", prod.getSize())
}
