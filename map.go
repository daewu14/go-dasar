package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Daewu",
		"address": "Yogyakarta",
	}

	person["hoby"] = "Workout"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["hoby"])

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Daewu"
	book["aduh"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "aduh")

	fmt.Println(book)
	fmt.Println(len(book))
}
