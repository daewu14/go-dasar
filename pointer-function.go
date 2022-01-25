package main

import "fmt"

type Address1 struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address1) {
	address.Country = "Indoneisa"
}

func main() {
	address := Address1{
		"Sleman",
		"DI Yogyakarta",
		"-",
	}

	ChangeAddressToIndonesia(&address)

	fmt.Println(address)
}
