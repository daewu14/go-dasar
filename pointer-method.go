package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Merried() {
	man.Name = "Mr. " + man.Name
}

func main() {
	man := Man{"Daewu"}
	man.Merried()

	fmt.Println(man)
}
