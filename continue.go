package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {

		// Maka jika i = 5 tidak ada eksekusi lanjut ke perulangan berikutnya
		if i == 5 {
			continue
		}

		fmt.Println("Perulangan ke:", i)

	}

}
