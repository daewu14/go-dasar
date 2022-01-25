package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"Novrmber",
		"Desember",
	}
	fmt.Println(months)

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// Merubah array maka slice juga berubah
	months[5] = "JUNI"
	fmt.Println(slice1)

	// Merubah slice maka yg di array juga berubah
	slice1[0] = "MEI"
	fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	// Jika array months sudah mlebihi batas capacity slice3 karena di append
	// maka akan menjadi sebuah variabel tidak mereferensikan ke array(months) jika ada perubahan di slice3
	// Jika kebalikannya maka akan merefrensikan perubahan ke array months
	var slice3 = append(slice2, "Daewu")
	fmt.Println(slice3)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Daewu"
	newSlice[1] = "Bintara"

	fmt.Println(newSlice)
	fmt.Println(cap(newSlice))
	fmt.Println(len(newSlice))

	copySlice := make([]string, len(newSlice), 5)

	copy(copySlice, newSlice)

	fmt.Println(copySlice)

	// Pembeda slice dan array sangat tipis
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
