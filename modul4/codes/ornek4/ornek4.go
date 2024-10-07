package main

import (
	"fmt"
)

// Uçabilen arayüz
type Flyer interface {
	Fly() string
}

// Su üzerinde hareket edebilen arayüz
type Swimmer interface {
	Swim() string
}

// Kuş ve balık arayüzü: Uçma ve yüzme yeteneklerini içerir
type BirdFish interface {
	Flyer
	Swimmer
}

// BirdFish türü
type Duck struct{}

func (d Duck) Fly() string {
	return "The duck flies!"
}

func (d Duck) Swim() string {
	return "The duck swims!"
}

func main() {
	var birdFish BirdFish

	// Duck türünü kullanma
	birdFish = Duck{}
	fmt.Println(birdFish.Fly())  // Çıktı: The duck flies!
	fmt.Println(birdFish.Swim()) // Çıktı: The duck swims!
}
