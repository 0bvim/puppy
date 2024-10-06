package puppy

import (
	"github.com/0bvim/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func FakeBark() string {
	return "Meow!"
}

func FakeBarks() string {
	return "Meow! Meow! Meow!"
}
