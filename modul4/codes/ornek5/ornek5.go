package main

import (
	"fmt"
)

// Arayüzler
type Speaker interface {
	Speak() string
}

type Mover interface {
	Move() string
}

type Animal interface {
	Speaker
	Mover
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "The dog runs!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Move() string {
	return "The cat jumps!"
}

func main() {
	var animal Animal

	animal = Dog{}
	speakAndMove(animal)

	animal = Cat{}
	speakAndMove(animal)
}

// Type assertion ile farklı türleri kontrol etme
func speakAndMove(a Animal) {
	fmt.Println(a.Speak())
	fmt.Println(a.Move())

	// Type assertion ile Dog kontrolü
	if dog, ok := a.(Dog); ok {
		fmt.Println("This is a dog:", dog.Speak())
	}

	// Type assertion ile Cat kontrolü
	if cat, ok := a.(Cat); ok {
		fmt.Println("This is a cat:", cat.Speak())
	}
}
