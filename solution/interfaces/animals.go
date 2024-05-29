package interfaces

import "fmt"

type Animal interface {
	Speak() string
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Wouf!"
}

func MakeSomeNoise(animals []Animal) {
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
