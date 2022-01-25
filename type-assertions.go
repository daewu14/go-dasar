package main

import "fmt"

func random() interface{} {
	return 100
}

func main() {
	result := random()
	//stringRes := result.(string)
	//fmt.Println(stringRes)

	//intRes := result.(int)
	//println(intRes) // Panic

	//cara agar tidak panic
	switch value := result.(type) {
	case int:
		fmt.Println("Value is", value, "Int")
	case string:
		fmt.Println("Value is", value, "String")
	}
}
