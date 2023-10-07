package main

import "fmt"

type Product interface {
	Use() string
}

type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Product A is being used"
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Product B is being used"
}

type Factory interface {
	CreateProduct() Product
}

type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	factoryA := &ConcreteFactoryA{}
	productA := factoryA.CreateProduct()
	fmt.Println(productA.Use())

	factoryB := &ConcreteFactoryB{}
	productB := factoryB.CreateProduct()
	fmt.Println(productB.Use())
}
