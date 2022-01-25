package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Daewu"
	names[1] = "Gus"
	names[2] = "Bintara"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	values := [3]int8{
		10,
		20,
		30,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))
}
