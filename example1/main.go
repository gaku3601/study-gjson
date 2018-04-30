package main

import (
	"fmt"
	"reflect"

	"github.com/tidwall/gjson"
)

func main() {
	json := `
	{
		"name":"gaku",
		"age": 26,
		"isMan": true
	}
	`
	name := gjson.Get(json, "name")
	age := gjson.Get(json, "age")
	isMan := gjson.Get(json, "isMan")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMan)
	fmt.Println(reflect.TypeOf(name.String()))
	fmt.Println(reflect.TypeOf(age.Int()))
	fmt.Println(reflect.TypeOf(isMan.Bool()))
}
