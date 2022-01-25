package main

import "fmt"

func main() {
	runApp()
}

func logging() {
	fmt.Println("Logging")
}

func runApp() {
	defer logging() // DIpanggil setelah fungsi selesai di eksekusi saat gak error ataupun error
	fmt.Println("Run Application")
}
