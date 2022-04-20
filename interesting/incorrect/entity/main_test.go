package entity_test

import (
	"testing"

	"github.com/tjgurwara99/goteach/interesting/incorrect/entity"
)

func TestSpeakers(t *testing.T) {
	speakers := []entity.Speaker{
		&entity.Cat{entity.Animal{"Mimi", 3}, 5},
		&entity.Dog{entity.Animal{"Barky", 4}, 2},
		&entity.Animal{"Buddy", 1},
	}
	for _, speaker := range speakers {
		t.Log(speaker.Speak())
	}
}
