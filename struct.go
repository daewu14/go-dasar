package main

import "fmt"

type User struct {
	Name, Email string
	Id          int
}

// Struct Function
func (user User) sayHi(name string) {
	fmt.Println("Hello", name, "my name is", user.Name)
}

func main() {

	var user User
	user = User{
		Id:    1,
		Email: "daewu@mail.com",
		Name:  "Daewu Gus Bintara",
	}

	user.Name = "Daewu Aja"
	user.sayHi("Joko")

	fmt.Println(user)
}
