package puppy

import {
	"github.com/matif607/dog"
}

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woo! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Bark())
}