package main

import (
	"fmt"
	"reflect"
)

func printMeta(obj any) {
	// pair: <value, type>
	t := reflect.TypeOf(obj)
	name := t.Name()
	kind := t.Kind()
	value := reflect.ValueOf(obj)
	fmt.Printf("Type: %s Type.Name: %s Kind: %s Value: %v\n", t, name, kind, value)
}

type handler func(int, int) int

func main() {

	var intVar int64 = 10
	stringVar := "hello"
	type book struct {
		name  string
		pages int
	}
	sum := func(a, b int) int {
		return a + b
	}
	var sub handler = func(a, b int) int {
		return a - b
	}
	sli := make([]int, 5)

	printMeta(intVar)
	printMeta(stringVar)
	printMeta(book{
		name:  "harry potter",
		pages: 500,
	})
	printMeta(sum)
	printMeta(sub)
	printMeta(sli)
}
