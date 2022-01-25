package main

import "fmt"

type Model interface{}

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func MyPrint(data Model) {
	fmt.Println(data)
}

func main() {
	data := Ups(5)

	fmt.Println(data)

	MyPrint("Okee")
}
