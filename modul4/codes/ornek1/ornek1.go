package main

import (
	"fmt"
)

// Arayüz tanımı
type Animal interface {
	Speak() string
}

// Arayüzü uygulayan bir tür
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	var animal Animal

	// Dog türünü kullanma
	animal = Dog{}
	fmt.Println(animal.Speak()) // Çıktı: Woof!

	// Cat türünü kullanma
	animal = Cat{}
	fmt.Println(animal.Speak()) // Çıktı: Meow!
}