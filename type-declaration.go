package main

import "fmt"

func main() {

	// Alias tipe data
	type Name string
	type Age int8

	var myName Name = "Daewu"
	var myAge Age = 25

	fmt.Println(myName)
	fmt.Println(myAge)
}
