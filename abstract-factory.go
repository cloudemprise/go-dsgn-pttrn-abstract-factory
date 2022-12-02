package main

/* abstract factory */

import "fmt"

// define an abstract factory that makes abstract products
type iAbstractFactory interface {
	makeProductA() iAbstractProductA
	makeProductB() iAbstractProductB
}

// free-standing abstract factory constructor
func newAbstractFactory(factoryName string) (iAbstractFactory, error) {
	var newFactory = new(concreteFactory)
	newFactory.name = factoryName
	return newFactory, nil
}

// define an abstract product (A) with some methods
type iAbstractProductA interface {
	setPropertyA(attribute string)
	getPropertyA() string
	setSizeOfA(size int)
	getSizeOfA() int
}

// abstract product (A) method constructor
func (*concreteFactory) makeProductA() iAbstractProductA {
	var produceA = new(concreteProductA)
	return produceA
}

// define an abstract product (B) with some methods
type iAbstractProductB interface {
	setPropertyB(attribute string)
	getPropertyB() string
	setSizeOfB(size int)
	getSizeOfB() int
}

// abstract product (B) method constructor
func (*concreteFactory) makeProductB() iAbstractProductB {
	var produceB = new(concreteProductB)
	return produceB
}

// define a concrete factory type
type concreteFactory struct {
	name string
}

// define a concrete product (A) type
type concreteProductA struct {
	attribute string
	size      int
}

// define a concrete product (B) type
type concreteProductB struct {
	attribute string
	size      int
}

// concrete productA implementations
// ---------------------------------

func (item *concreteProductA) setPropertyA(attribute string) {
	item.attribute = attribute
}

func (item *concreteProductA) getPropertyA() string {
	return item.attribute
}

func (item *concreteProductA) setSizeOfA(size int) {
	item.size = size
}

func (s *concreteProductA) getSizeOfA() int {
	return s.size
}

// concrete productB implementations
// ---------------------------------

func (item *concreteProductB) setPropertyB(attribute string) {
	item.attribute = attribute
}

func (item *concreteProductB) getPropertyB() string {
	return item.attribute
}

func (item *concreteProductB) setSizeOfB(size int) {
	item.size = size
}

func (item *concreteProductB) getSizeOfB() int {
	return item.size
}

func main() {
	// create a new abstract factory
	myFactory, _ := newAbstractFactory("myAwesomeBrand")

	// create productA from abstract factory
	myProductA := myFactory.makeProductA()
	makeAbstractProductA(myProductA)

	// create productB from abstract factory
	myProductB := myFactory.makeProductB()
	makeAbstractProductB(myProductB)
}

func makeAbstractProductA(prodA iAbstractProductA) {

	prodA.setPropertyA("Set product A attribute here.")
	prodA.setSizeOfA(69)
	fmt.Printf("Product A Property: %s\n", prodA.getPropertyA())
	fmt.Printf("Product A Size     : %d\n", prodA.getSizeOfA())
}

func makeAbstractProductB(prodB iAbstractProductB) {

	prodB.setPropertyB("Set product B attribute here.")
	prodB.setSizeOfB(13)
	fmt.Printf("Product B Property: %s\n", prodB.getPropertyB())
	fmt.Printf("Product B Size     : %d\n", prodB.getSizeOfB())
}
