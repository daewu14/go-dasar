package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHelo(hasName HasName) {
	fmt.Println("Helo", hasName.GetName())
}

type Person struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}

func main() {

	var person Person
	person = Person{
		Name: "Daewu",
	}
	fmt.Println(person.GetName())

	person.Name = "Daewu 2"
	fmt.Println(person.GetName())
}
