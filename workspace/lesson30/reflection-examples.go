package main

import (
	"log"
	"reflect"
	"sync"
)

type User struct {
	name string
	age  int
}

func (u User) PrintName() {
	log.Println(u.name)
}

func (u User) PrintAge() {
	log.Println(u.age)
}

type Aggregator func(int, int) int

var (
	user = User{
		name: "heli",
		age:  24,
	}

	add Aggregator = func(a, b int) int {
		return a + b
	}
	sub Aggregator = func(a, b int) int {
		return a - b
	}

	wg sync.WaitGroup
)

func inspect(variable interface{}) {
	t := reflect.TypeOf(variable)
	v := reflect.ValueOf(variable)

	//如果是struct类型
	if t.Kind() == reflect.Struct {
		// print its fields
		log.Printf("--------- fields %d ----------\n", t.NumField())
		for idx := 0; idx < t.NumField(); idx++ {
			fieldType := t.Field(idx)
			fieldVal := v.Field(idx)
			log.Printf("\t %v = %v\n", fieldType.Name, fieldVal)
		}

		// iterate its methods
		log.Printf("--------- methods %d ----------\n", t.NumMethod())
		for idx := 0; idx < t.NumMethod(); idx++ {
			methodType := t.Method(idx)
			log.Printf("\t method_name=%s input_num=%d, output_num=%d\n",
				methodType.Name,
				methodType.Type.NumIn(),
				methodType.Type.NumOut())
		}
	}

	//如果是函数类型
	if t.Kind() == reflect.Func {
		log.Printf("%s function has %d inputs and %d outputs\n",
			t.Name(),
			t.NumIn(),
			t.NumOut())
	}
	log.Printf("\n\n")
}

func main() {

	wg.Add(3)

	go func() {
		defer wg.Done()
		inspect(user)
	}()
	go func() {
		defer wg.Done()
		inspect(add)
	}()
	go func() {
		defer wg.Done()
		inspect(sub)
	}()

	wg.Wait()

}
