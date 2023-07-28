package main

import (
	"fmt"
	"reflect"
	"time"
)

type Student struct {
	name string
	age  int
}

func (s *Student) DoHomework(number int) {
	fmt.Printf("%s is doing homework %d,age is now %d\n", s.name, number, s.age)
	s.age--
	time.Sleep(3 * time.Second)
}

func (s *Student) Sleep() {
	fmt.Printf("%s is sleeping\n", s.name)
}

func main() {
	// use reflect to invoke the DoHomework of a student
	s := Student{
		name: "heli",
		age:  24,
	}
	v := reflect.ValueOf(&s)

	t := v.Type()
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("struct %d method, name:%s type:%v\n", i, method.Name, method.Type)
	}

	methodV := v.MethodByName("DoHomework")
	if methodV.IsValid() {
		in := []reflect.Value{reflect.ValueOf(55)}
		methodV.Call(in)
	}

	fmt.Printf("age is now %d\n", s.age)

	methodV2 := v.MethodByName("Sleep")
	if methodV2.IsValid() {
		in := []reflect.Value{}
		methodV2.Call(in)
	}
}
