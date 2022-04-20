package entity

import "strconv"

type Animal struct {
	Name string
	Age  int
}

func (a *Animal) Speak() string {
	return "Hello, my name is " +
		a.Name +
		". I'm an animal and I'm " +
		strconv.Itoa(a.Age) +
		" years old."
}

func (a *Animal) Category() string {
	return "animal"
}
