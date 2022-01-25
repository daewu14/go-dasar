package main

import "fmt"

func main() {
	firstName, lastName := getFullName()

	fmt.Println(firstName)
	fmt.Println(lastName)

	// Jika lastName tidak ingin di gunakan
	namaDepan, _ := getFullName()

	fmt.Println(namaDepan)
}

func getFullName() (string, string) {
	return "Daewu", "Bintara"
}
