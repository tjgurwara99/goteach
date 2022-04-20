package entity

import "strconv"

type Cat struct {
	Animal
	ClimbHeight int
}

func (c *Cat) Speak() string {
	return "Meow! I can climb " + strconv.Itoa(c.ClimbHeight) + " feet height."
}
