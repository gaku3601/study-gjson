package main

import (
	"fmt"
	"reflect"

	"github.com/tidwall/gjson"
)

func main() {
	json := `
	[{
		"name":"gaku",
		"age": 26,
		"isMan": true,
		"children":[{
			"name": "masako",
			"age": 2,
			"isMan": false
		},{
			"name": "masao",
			"age": 15,
			"isMan": true
		}]
	},{
		"name":"gakuko",
		"age": 20,
		"isMan": false,
		"children":[{
			"name": "masako2",
			"age": 2,
			"isMan": false
		},{
			"name": "masao2",
			"age": 15,
			"isMan": true
		}]
	}]
	`
	res := gjson.Parse(json)
	for _, v := range res.Array() {
		name := v.Get("name")
		age := v.Get("age")
		isMan := v.Get("isMan")
		fmt.Println(name)
		fmt.Println(age)
		fmt.Println(isMan)
		fmt.Println(reflect.TypeOf(name.String()))
		fmt.Println(reflect.TypeOf(age.Int()))
		fmt.Println(reflect.TypeOf(isMan.Bool()))
		fmt.Println("--------------------")
	}
}
