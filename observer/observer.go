package main

import "fmt"

func main() {
	s := Subject{}

	var binary Binary = Binary{}
	var octal Octal = Octal{}
	binary.BinaryListener(&s)
	octal.OctalListener(&s)

	s.setState(10)
	s.setState(15)
}

type Observer interface {
	update()
}

type Subject struct {
	state     int
	observers []Observer
}

func (s Subject) getState() int {
	return s.state
}

func (s *Subject) attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s Subject) notifyAllService() {
	for _, ser := range s.observers {
		ser.update()
	}
}

func (s *Subject) setState(state int) {
	s.state = state
	s.notifyAllService()
}

type Binary struct {
	subject *Subject
}

func (b Binary) BinaryListener(sub *Subject) {
	b.subject = sub
	sub.attach(b)
}

func (b Binary) update() {
	fmt.Println("Binary ", b.subject.getState())
}

type Octal struct {
	subject *Subject
}

func (b Octal) OctalListener(sub *Subject) {
	b.subject = sub
	sub.attach(b)
}

func (b Octal) update() {
	fmt.Println("Octal ", b.subject.getState())
}
