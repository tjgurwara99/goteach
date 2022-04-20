package entity

import "strconv"

type Dog struct {
	Animal
	Friends int
}

func (d *Dog) Speak() string {
	return "Woof! I have " + strconv.Itoa(d.Friends) + " friends."
}
