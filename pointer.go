package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Sleman", "DI Yogyakarta", "Indonesia"}
	address2 := address1

	address2.City = "Yogya"

	fmt.Println(address1) // Tidak berubah
	fmt.Println(address2)

	address3 := &address2
	address3.City = "Yogya"

	fmt.Println(address2) // Berubah
	fmt.Println(address3)

	// Melakukan instance ulang tanpa mempengaruhi semua
	address4 := Address{"Malang", "Jawa timur", "Indonesia"}

	// Melakukan instance ulang dengan mempengaruhi semua
	address5 := address3
	*address5 = Address{"Kulonprogo", "DI Yogyakarta", "Indonesia"}

	fmt.Println(address5)
	fmt.Println(address4)
	fmt.Println(address2)
	fmt.Println(address3)

	address6 := new(Address)
	address6.Country = "Indonesia"
	fmt.Println(address6)
}
