package main

/* abstract factory pattern */

import "fmt"

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : An Abstract Factory Interface

type iAbstractFactory interface {
	makeProductA() iAbstractProductA
	makeProductB() iAbstractProductB
}

// END : An Abstract Factory
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Concrete Factory Function

func createFactory(factory string) (iAbstractFactory, error) {
	if factory == "customFactory1" {
		return &concreteFactory1{}, nil
	}

	if factory == "customFactory2" {
		return &concreteFactory2{}, nil
	}

	return nil, fmt.Errorf("wrong factory type selected")
}

// END : Concrete Factory Function
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Concrete product constructors 1

// concrete factory type
type concreteFactory1 struct {
}

// method implementations

func (a *concreteFactory1) makeProductA() iAbstractProductA {
	prod := new(concreteProductA)
	prod.setBehaviourA("defaultA1")
	prod.setFeatureA(0)
	return prod
}

func (a *concreteFactory1) makeProductB() iAbstractProductB {
	prod := new(concreteProductB)
	prod.setBehaviourB("defaultB1")
	prod.setFeatureB(0)
	return prod
}

// END : Concrete product constructors 1
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Concrete product constructors 2

// concrete factory type
type concreteFactory2 struct {
}

// method implementations

func (a *concreteFactory2) makeProductA() iAbstractProductA {
	prod := new(concreteProductA)
	prod.setBehaviourA("defaultA2")
	prod.setFeatureA(0)
	return prod
}

func (a *concreteFactory2) makeProductB() iAbstractProductB {
	prod := new(concreteProductB)
	prod.setBehaviourB("defaultB2")
	prod.setFeatureB(0)
	return prod
}

// END : Concrete product constructors 2
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Abstract Product A

// product behaviours
type iAbstractProductA interface {
	setBehaviourA(behaviour string)
	getBehaviourA() string
	setFeatureA(feature int)
	getFeatureA() int
}

// END : Abstract Product A
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Concrete Product A Definition

type concreteProductA struct {
	behaviourA string
	featureA   int
}

// method implementations

func (p *concreteProductA) setBehaviourA(behaviour string) {
	p.behaviourA = behaviour
}

func (p *concreteProductA) getBehaviourA() string {
	return p.behaviourA
}

func (p *concreteProductA) setFeatureA(feature int) {
	p.featureA = feature
}

func (p *concreteProductA) getFeatureA() int {
	return p.featureA
}

// END : Concrete Product A Definition
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Abstract Product B

// product behaviours
type iAbstractProductB interface {
	setBehaviourB(behaviour string)
	getBehaviourB() string
	setFeatureB(feature int)
	getFeatureB() int
}

// END : Abstract Product B
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Concrete Product B Definition

type concreteProductB struct {
	behaviourB string
	featureB   int
}

// method implementations

func (p *concreteProductB) setBehaviourB(behaviour string) {
	p.behaviourB = behaviour
}

func (p *concreteProductB) getBehaviourB() string {
	return p.behaviourB
}

func (p *concreteProductB) setFeatureB(feature int) {
	p.featureB = feature
}

func (p *concreteProductB) getFeatureB() int {
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
	myFactory1, _ := createFactory("customFactory1")

	// create a new productA
	myProductA1 := myFactory1.makeProductA()
	// specify new productA characteristics
	myProductA1.setBehaviourA("myProductA1")
	myProductA1.setFeatureA(12)

	// print
	printProductA(myProductA1)

	// create a new productB
	myProductB1 := myFactory1.makeProductB()
	// specify new productB characteristics
	myProductB1.setBehaviourB("myProductB1")
	myProductB1.setFeatureB(34)

	// print
	printProductB(myProductB1)

	// ---------------------------
	// create a different factory
	myFactory2, _ := createFactory("customFactory2")

	// create a new productA
	myProductA2 := myFactory2.makeProductA()
	// specify new productA characteristics
	myProductA2.setBehaviourA("myProductA2")
	myProductA2.setFeatureA(56)

	// print
	printProductA(myProductA2)

	// create a new productB
	myProductB2 := myFactory2.makeProductB()
	// specify new productB characteristics
	myProductB2.setBehaviourB("myProductB2")
	myProductB2.setFeatureB(78)

	// print
	printProductB(myProductB2)

}

// Test helpers

func printProductA(s iAbstractProductA) {
	fmt.Printf("BehaviourA : %s", s.getBehaviourA())
	fmt.Println()
	fmt.Printf("FeatureA   : %d", s.getFeatureA())
	fmt.Println()
}

func printProductB(s iAbstractProductB) {
	fmt.Printf("BehaviourB : %s", s.getBehaviourB())
	fmt.Println()
	fmt.Printf("FeatureB   : %d", s.getFeatureB())
	fmt.Println()
}

// END : Client Code
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
