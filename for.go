package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke :", counter)
	}

	slice := []string{
		"Daewu",
		"Gus",
		"Bintara",
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println("Index ke :", i, " adalah :", slice[i])
	}

	// For range
	for i, value := range slice {
		fmt.Println("Index", i, "adalah :", value)
	}

	// For range
	// jika i tidak mau digunakan
	for _, value := range slice {
		fmt.Println(value)
	}

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Daewu"
	book["aduh"] = "Salah"

	for key, value := range book {
		fmt.Println(key, ":", value)
	}

}
