package main

import (
	"fmt"
)

// Konuşma yeteneği
type Speaker interface {
	Speak() string
}

// Hareket etme yeteneği
type Mover interface {
	Move() string
}

// Hayvan arayüzü: Hem konuşma hem hareket etme yeteneğini içerir
type Animal interface {
	Speaker
	Mover
}

// Dog türü
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "The dog runs!"
}

// Cat türü
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Move() string {
	return "The cat jumps!"
}

func main() {
	var animal Animal

	// Dog türünü kullanma
	animal = Dog{}
	fmt.Println(animal.Speak()) // Çıktı: Woof!
	fmt.Println(animal.Move())  // Çıktı: The dog runs!

	// Cat türünü kullanma
	animal = Cat{}
	fmt.Println(animal.Speak()) // Çıktı: Meow!
	fmt.Println(animal.Move())  // Çıktı: The cat jumps!
}
