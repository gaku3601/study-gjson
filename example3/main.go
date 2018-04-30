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
	}
	`
	result := gjson.Get(json, "people.#.children.1.name")
	result.ForEach(func(key, value gjson.Result) bool {
		fmt.Println(value.String())
		fmt.Println(reflect.TypeOf(value.String()))
		return true // keep iterating
	})
}
