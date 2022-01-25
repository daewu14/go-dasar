package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (hasil int, err error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}
	hasil = nilai / pembagi
	return hasil, nil
}

func main() {

	hasil, err := Pembagi(10, 2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(hasil)
	}

}
