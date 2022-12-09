package main

/* abstract factory pattern */

import "fmt"

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : An Abstract Factory

// an abstract factory that makes abstract products
type iAbstractFactory interface {
	// abstract pointers to a concrete objects
	makeProductA() iAbstractProduct
	makeProductB() iAbstractProduct
}

// END : An Abstract Factory
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Concrete Factory Function

func createFactory() iAbstractFactory {
	var factory = new(concreteFactory)
	return factory

}

// END : Concrete Factory Function
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Concrete Factory Type

// a concrete factory type
type concreteFactory struct {
	//specialAttribute string // factory distinction
}

// concreteFactory method implementation of iAbstractFactory interface.
func (*concreteFactory) makeProductA() iAbstractProduct {
	var p = new(concreteProductA)
	/* 	p.behaviourA = "default"
	   	p.featureA = 0 */
	return p
}

// concreteFactory method implementation of iAbstractFactory interface.
func (*concreteFactory) makeProductB() iAbstractProduct {
	var p = new(concreteProductB)
	/* 	p.behaviourA = "default"
	   	p.featureA = 0 */
	return p
}

// END : Concrete Factory 1 Type
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Abstract Product Interface

// define product behaviours
type iAbstractProduct interface {
	setBehaviour(behaviour string)
	getBehaviour() string
	setFeature(feature int)
	getFeature() int
}

// END : Abstract Product Interface
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Concrete Product A Definition

type concreteProductA struct {
	// products defines similar attributes
	behaviourA string
	featureA   int
}

// Concrete Product A Implementation
// ---------------------------------

func (p *concreteProductA) setBehaviour(behaviour string) {
	p.behaviourA = behaviour
}

func (p *concreteProductA) getBehaviour() string {
	return p.behaviourA
}

func (p *concreteProductA) setFeature(feature int) {
	p.featureA = feature
}

func (p *concreteProductA) getFeature() int {
	return p.featureA
}

// END : Concrete Product A Definition
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Concrete Product B Definition

type concreteProductB struct {
	// products defines similar attributes
	behaviourB string
	featureB   int
}

// Concrete Product B Implementation
// ---------------------------------

func (p *concreteProductB) setBehaviour(behaviour string) {
	p.behaviourB = behaviour
}

func (p *concreteProductB) getBehaviour() string {
	return p.behaviourB
}

func (p *concreteProductB) setFeature(feature int) {
	p.featureB = feature
}

func (p *concreteProductB) getFeature() int {
	return p.featureB
}

// END : Concrete Product B Definition
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Client Code

func main() {

	// ---------------------------
	// create a new factory
	myFactory1 := createFactory()

	// create a new productA
	myProductA1 := myFactory1.makeProductA()
	// specify new productA characteristics
	myProductA1.setBehaviour("myProductA1")
	myProductA1.setFeature(12)

	// print
	getProduct(myProductA1)

	// create a new productB
	myProductB1 := myFactory1.makeProductB()
	// specify new productB characteristics
	myProductB1.setBehaviour("myProductB1")
	myProductB1.setFeature(34)

	// print
	getProduct(myProductB1)

	// ---------------------------
	// create a different factory
	myFactory2 := createFactory()

	// create a new productA
	myProductA2 := myFactory2.makeProductA()
	// specify new productA characteristics
	myProductA2.setBehaviour("myProductA2")
	myProductA1.setFeature(56)

	// print
	getProduct(myProductA1)

	// create a new productB
	myProductB2 := myFactory2.makeProductB()
	// specify new productB characteristics
	myProductB2.setBehaviour("myProductB2")
	myProductB2.setFeature(78)

	// print
	getProduct(myProductB2)

}

// Test helper
func getProduct(p iAbstractProduct) {
	fmt.Println(p.getBehaviour())
	fmt.Println(p.getFeature())
}

// END : Client Code
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
