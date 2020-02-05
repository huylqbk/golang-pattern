package main

import "fmt"

func main() {
	var phone Phone = FactoryInit(ApplePhone)
	phone.showBrand()
}

type TypePhone uint8

const (
	SamsungPhone TypePhone = iota
	ApplePhone
	XiaomiPhone
)

type Phone interface {
	showBrand()
}

type Samsung struct{}

func (s Samsung) showBrand() {
	fmt.Println("Samsung")
}

type Apple struct{}

func (a Apple) showBrand() {
	fmt.Println("Apple")
}

type Xiaomi struct{}

func (x Xiaomi) showBrand() {
	fmt.Println("Xiaomi")
}

func FactoryInit(t TypePhone) Phone {
	if t == SamsungPhone {
		return Samsung{}
	}
	if t == ApplePhone {
		return Apple{}
	}
	if t == XiaomiPhone {
		return Xiaomi{}
	}
	return nil
}
