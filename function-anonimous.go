package main

import "fmt"

type Blaclist func(string) bool

func main() {

	blaclist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("anjing", blaclist)

	registerUser("Daewu", func(name string) bool {
		return name == "anjing"
	})
}

func registerUser(name string, blacklist Blaclist) {

	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}

}
