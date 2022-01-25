package main

import "fmt"

func main() {
	var result = 3 + 10
	fmt.Println(result)

	var (
		a = 5
		b = 10
		c = a * b
	)
	fmt.Println(c)

	c += 1
	b -= 1
	fmt.Println(a)
	fmt.Println(b)

	// Yg lain mirip dengan bhs pemrograman php, dll
}
