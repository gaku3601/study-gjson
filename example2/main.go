package main

import (
	"fmt"
	"reflect"

	"github.com/tidwall/gjson"
)

func main() {
	json := `
	{
		"people":[{
			"name":"gaku",
			"age": 26,
			"isMan": true
		},{
			"name":"gakuko",
			"age": 20,
			"isMan": false
		}]
	}
	`
	result := gjson.Get(json, "people")
	result.ForEach(func(key, value gjson.Result) bool {
		name := value.Get("name")
		age := value.Get("age")
		isMan := value.Get("isMan")
		fmt.Println(name)
		fmt.Println(age)
		fmt.Println(isMan)
		fmt.Println(reflect.TypeOf(name.String()))
		fmt.Println(reflect.TypeOf(age.Int()))
		fmt.Println(reflect.TypeOf(isMan.Bool()))
		return true // keep iterating
	})
}
