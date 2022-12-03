package main

/* abstract factory */

import "fmt"

// define an abstract factory that makes abstract products
type iAbstractFactory interface {
	// factory returns an abstract pointer to a concrete object
	makeProduct() iAbstractProduct
}

// a free-standing abstract factory constructor
func newAbstractFactory(factoryAttribute string) (iAbstractFactory, error) {
	var newFactory = new(concreteFactory)
	newFactory.specialAttribute = factoryAttribute
	return newFactory, nil
}

// define an abstract product with some behaviours
type iAbstractProduct interface {
	setBehaviour(behaviour string)
	getBehaviour() string
	setFeature(feature int)
	getFeature() int
}

// an abstract product constructor
func (*concreteFactory) makeProduct() iAbstractProduct {
	var produce = new(concreteProduct)
	return produce
}

// define a concrete factory type
type concreteFactory struct {
	specialAttribute string
}

// define a concrete product type
type concreteProduct struct {
	behaviour string
	feature   int
}

// concrete product implementations
// ---------------------------------

func (item *concreteProduct) setBehaviour(behaviour string) {
	item.behaviour = behaviour
}

func (item *concreteProduct) getBehaviour() string {
	return item.behaviour
}

func (item *concreteProduct) setFeature(feature int) {
	item.feature = feature
}

func (item *concreteProduct) getFeature() int {
	return item.feature
}

func main() {
	// create a new abstract factory
	myFactory, _ := newAbstractFactory("myAwesomeBrand")

	// create product from abstract factory
	myProduct := myFactory.makeProduct()
	makeAbstractProduct(myProduct)

}

// Test Case
func makeAbstractProduct(product iAbstractProduct) {

	product.setBehaviour("Special Behaviour")
	product.setFeature(69)
	fmt.Printf("Product Behaviour : %s\n", product.getBehaviour())
	fmt.Printf("Product Feature   : %d\n", product.getFeature())
}
