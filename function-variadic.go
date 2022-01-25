package main

import "fmt"

func main() {

	slice := []int{10, 20}

	total := sumAll(10, 10)

	fmt.Println("Total :", total)
	fmt.Println("Total Sum Slice :", sumAll(slice...))

}

func sumAll(numbers ...int) (total int) {
	total = 0
	for _, value := range numbers {
		total += value
	}
	return total
}
