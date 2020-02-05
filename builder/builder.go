package main

import "fmt"

func main() {
	people := newInstance().setName("A").setOld(10).setEmail("a@gmail.com").builder()
	people.toString()
}

type People struct {
	Name  string
	Old   int
	Email string
}

func (p *People) People(b *Builder) {
	p.Name = b.Name
	p.Old = b.Old
	p.Email = b.Email
}

func (p People) toString() {
	fmt.Println(p.Name, p.Old, p.Email)
}

type Builder struct {
	Name  string
	Old   int
	Email string
}

func newInstance() *Builder {
	return &Builder{}
}

func (b *Builder) setName(name string) *Builder {
	b.Name = name
	return b
}

func (b *Builder) setOld(old int) *Builder {
	b.Old = old
	return b
}

func (b *Builder) setEmail(email string) *Builder {
	b.Email = email
	return b
}

//builder return People object from Builder struct
func (b *Builder) builder() People {
	p := &People{}
	p.People(b)
	return *p
}
