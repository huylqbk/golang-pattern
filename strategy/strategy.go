package main

import "fmt"

func main() {
	c := Context{}

	c.Set(Substract{})
	fmt.Println(c.Exec(10, 5))
}

type Context struct {
	Operate Operator
}

func (c *Context) Set(o Operator) {
	c.Operate = o
}

func (c Context) Exec(v1, v2 int) int {
	return c.Operate.do(v1, v2)
}

type Operator interface {
	do(int, int) int
}

type Substract struct{}

func (s Substract) do(v1, v2 int) int {
	return v1 - v2
}

type Addition struct{}

func (a Addition) do(v1, v2 int) int {
	return v1 - v2
}

type Multiply struct{}

func (m Multiply) do(v1, v2 int) int {
	return v1 * v2
}
