package main

type human interface {
	speak()
	think()
	eat()
	sex() string
	breathe()
}

type animal interface {
	breathe()
	eat()
	isCarnivorous() bool
}

type men struct {
	name      string
	age       int
	height    float64
	weight    float64
	breathing bool
	eating    bool
	thinking  bool
}

type women struct {
	name      string
	age       int
	height    float64
	weight    float64
	breathing bool
	eating    bool
	thinking  bool
}

func (m *men) breathe() {
	m.breathing = true
}

func (m *men) eat() {
	m.eating = true
}

func (m *men) think() {
	m.thinking = true
}

func main() {

}
