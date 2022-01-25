package main

import "fmt"

type Filter func(string) string

func main() {
	sayHelloWithFilter("Daewu", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	sayHelloWithFilterDeclarative("Daewu", spamFilter)
	sayHelloWithFilterDeclarative("Anjing", spamFilter)
}

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Halo", filter(name))
}

func sayHelloWithFilterDeclarative(name string, filter Filter) {
	fmt.Println("Halo", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}
	return name
}
