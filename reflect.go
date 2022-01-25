package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func main() {

	sample := Sample{"Daewu"}
	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType.Name())
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
}
