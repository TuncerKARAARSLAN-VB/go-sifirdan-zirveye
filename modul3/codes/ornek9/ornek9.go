package main

import "fmt"

// Address adında bir struct tanımlama
type Address struct {
	City    string
	Country string
}

// Person adında bir struct tanımlama, Address içeren
type Person struct {
	Name    string
	Age     int
	Address Address // İç içe struct
}

func main() {
	// Struct kullanarak bir nesne oluşturma
	person1 := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			City:    "Istanbul",
			Country: "Türkiye",
		},
	}

	// Struct içindeki verilere erişim
	fmt.Println("İsim:", person1.Name)
	fmt.Println("Yaş:", person1.Age)
	fmt.Println("Şehir:", person1.Address.City)
	fmt.Println("Ülke:", person1.Address.Country)
}
