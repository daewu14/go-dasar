package main

import "fmt"

func main() {
	// Inisialisasi tipe data
	var name string

	fmt.Println(name)

	if name == "" {
		name = "Daewu Bintara"
	}
	fmt.Println(name)

	name = "Daewu Gus"
	fmt.Println(name)

	var friendName = "Budi"
	fmt.Println(friendName)

	// sama dengan var country = "Indonesia"
	country := "Indonesia"
	fmt.Println(country)

	country = "Indonesia 2"
	fmt.Println(country)

	// Multi variable
	var (
		firstName = "Daewu"
		lastName  = "Bintara"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
