package main

import (
	"fmt"
	"sync"
)

func main() {
	var db DB = New()
	db.Connect()
}

type DB interface {
	Connect()
}

var (
	one      sync.Once
	instance DB
)

type Mysql struct{}

func (m Mysql) Connect() {
	fmt.Println("Mysql Connected")
}

func New() DB {
	one.Do(
		func() {
			instance = Mysql{}
		})
	return instance
}
