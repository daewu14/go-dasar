package main

import "fmt"

func main() {

	helo := getHello()
	heloTo := getHelloTo("Daewu")

	fmt.Println(helo)
	fmt.Println(heloTo)
}

func getHello() string {
	return "Hello"
}

func getHelloTo(name string) string {
	return "Hello = " + name
}
