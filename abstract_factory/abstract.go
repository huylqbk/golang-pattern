package main

import (
	"fmt"
)

func main() {
	var factory AbstractFactory
	factory = getFactory(PlasticType)

	var chair Furniture
	chair = factory.createChair()
	var table Furniture
	table = factory.createTable()
	chair.show()
	table.show()

	fmt.Println("-----------------")

	factory = getFactory(WoodType)
	chair = factory.createChair()
	table = factory.createTable()
	chair.show()
	table.show()

}

type FurnitureType uint

const (
	PlasticType FurnitureType = iota
	WoodType
)

type AbstractFactory interface {
	createChair() Furniture
	createTable() Furniture
}

type Furniture interface {
	show()
}

type WoodFactory struct{}

type PlasticFactory struct{}

func (w WoodFactory) createChair() Furniture {
	return WoodChair{}
}

func (w WoodFactory) createTable() Furniture {
	return WoodTable{}
}

func (p PlasticFactory) createChair() Furniture {
	return PlasticChair{}
}

func (p PlasticFactory) createTable() Furniture {
	return PlasticTable{}
}

func getFactory(fType FurnitureType) AbstractFactory {
	switch fType {
	case PlasticType:
		return PlasticFactory{}
	case WoodType:
		return WoodFactory{}
	default:
		return nil
	}
}

type PlasticChair struct{}
type WoodChair struct{}
type PlasticTable struct{}
type WoodTable struct{}

func (pc PlasticChair) show() {
	fmt.Println("PlasticChair")
}

func (wc WoodChair) show() {
	fmt.Println("WoodChair")
}

func (pt PlasticTable) show() {
	fmt.Println("PlasticTable")
}

func (wt WoodTable) show() {
	fmt.Println("WoodTable")
}
