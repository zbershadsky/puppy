package puppy

import (
	"github.com/zbershadsky/dog"
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

func HeyFromV1() string {
	return "It's v1.0.0"
}

func HeyFromV11() string {
	return "It's v1.1.0"
}
