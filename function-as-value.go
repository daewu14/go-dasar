package main

import "fmt"

func main() {
	saygoodby := getGoodBy

	goodByDaewu := saygoodby("Daewu")

	fmt.Println(goodByDaewu)
}

func getGoodBy(name string) string {
	return "Good By " + name
}
