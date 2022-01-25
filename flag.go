package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "local", "Put your host")
	var user *string = flag.String("user", "root", "Put your username")
	var password *string = flag.String("password", "", "Put your username")

	flag.Parse()

	fmt.Println("Host :", *host)
	fmt.Println("User :", *user)
	fmt.Println("password :", *password)
}
