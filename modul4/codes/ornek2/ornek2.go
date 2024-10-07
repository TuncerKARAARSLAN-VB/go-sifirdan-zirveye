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

func main() {
	var animal Animal = Dog{}

	// Type assertion
	dog, ok := animal.(Dog)
	if ok {
		fmt.Println("Animal is a Dog:", dog.Speak())
	} else {
		fmt.Println("Animal is not a Dog")
	}
}
