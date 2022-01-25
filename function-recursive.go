package main

import "fmt"

func main() {

	loop := factorialLoop(5)
	fmt.Println(loop)

	revursive := factorialRecursive(5)
	fmt.Println(revursive)

}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}
