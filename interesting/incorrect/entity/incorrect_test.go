package entity_test

import (
	"testing"

	"github.com/tjgurwara99/goteach/interesting/incorrect/entity"
)

// TestAnimalsRange is a test to demonstrate that this test is not correct.
// func TestAnimalsRange(t *testing.T) {
// 	animals := []Animal{
// 		&Cat{Animal{"Mimi", 3}, 5},
// 		&Dog{Animal{"Barky", 4}, 2},
// 	}

// 	for _, animal := range animals {
// 		t.Log(animal.Speak())
// 	}
// }

func TestAnimalCategory(t *testing.T) {
	animals := entity.Animal{
		Name: "Buddy",
		Age:  1,
	}
	t.Log(animals.Category())

	cats := entity.Cat{
		Animal: entity.Animal{
			Name: "Mimi",
			Age:  3,
		},
		ClimbHeight: 5,
	}
	t.Log(cats.Category()) // why? if its not inheritence?

	dogs := entity.Dog{
		Animal: entity.Animal{
			Name: "Barky",
			Age:  4,
		},
		Friends: 2,
	}
	t.Log(dogs.Category()) // why? if its not inheritence?
}

func TestAnimalCategory2(t *testing.T) {
	animal := entity.Animal{
		Name: "Buddy",
		Age:  1,
	}
	t.Log(animal.Category())

	cat := entity.Cat{
		Animal: entity.Animal{
			Name: "Mimi",
			Age:  3,
		},
		ClimbHeight: 5,
	}
	t.Log(cat.Animal.Category())

	dog := entity.Dog{
		Animal: entity.Animal{
			Name: "Barky",
			Age:  4,
		},
		Friends: 2,
	}
	t.Log(dog.Animal.Category())
}

// This is done with something called Method Forwarding.
// essentially the above means that the methods in the animal got forwarded.
// For example
// func (c *Cat) Category() string {
// 	return c.Animal.Category()
// }
// This is done by compiler as a syntactic sugar. Personally, I argue that
// this is not a good practice for large scale programs.
