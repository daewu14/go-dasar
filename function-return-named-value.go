package main

import "fmt"

func main() {
	fName, lName, age := getFullNameWithAge()

	fmt.Println(fName)
	fmt.Println(lName)
	fmt.Println(age)
}

func getFullNameWithAge() (firstName string, lastName string, age int) {
	firstName = "Daewu"
	lastName = "Bintara"
	age = 25
	return
}
