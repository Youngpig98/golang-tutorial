package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string
	Age  int32
}

func main() {
	// FIRST EXAMPLE SHOWING CONVERT REFLECT.VALUE TO FLOAT
	floatVar := 3.14
	v := reflect.ValueOf(floatVar)

	fmt.Printf("Is floatVar canSet: %v , canAddr:%v", v.CanSet(), v.CanAddr())
	vPtr := reflect.ValueOf(&floatVar)
	fmt.Printf("Is floatVarPtr canSet: %v , canAddr:%v", vPtr.CanSet(), vPtr.CanAddr())

	newFloat := v.Interface().(float64)
	fmt.Println(newFloat + 1.2)

	stringVar := "HelloWorld"
	v = reflect.ValueOf(stringVar)
	t := v.Type()
	fmt.Printf("Type: %s Type.Name: %s Kind: %s\n", t, t.Name(), t.Kind())
	newStr := v.Interface().(string) + ",young!"
	fmt.Println(newStr)

	// second example showing convert Reflect.Value to slice
	sliceVar := make([]int, 5)
	v = reflect.ValueOf(sliceVar)
	v = reflect.Append(v, reflect.ValueOf(2))
	newSlice := v.Interface().([]int)
	newSlice = append(newSlice, 4)
	fmt.Println(newSlice)

	// third example showing convert Reflect.Value to student
	stuPtr := reflect.New(reflect.TypeOf(student{
		Name: "wangyuyang",
		Age:  24,
	}))
	stu := stuPtr.Elem() //取地址，相当于*
	nameField := stu.FieldByName("Name")
	if nameField.IsValid() {
		if nameField.CanSet() {
			nameField.SetString("chong")
		}
		realStudent := stu.Interface().(student)
		fmt.Println(realStudent)
	}

}
