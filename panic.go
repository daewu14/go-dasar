package main

import "fmt"

func main() {
	run(true)
}

func run(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan pesan :", message)
	}
	fmt.Println("End App")
}
