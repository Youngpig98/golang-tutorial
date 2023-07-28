package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("The fault is ", r)
		}
	}()
	fVar := 3.14
	v := reflect.ValueOf(fVar)
	fmt.Printf("is float canSet: %v canAddr %v\n", v.CanSet(), v.CanAddr())
	//v.SetFloat(221232.12)
	//panic("Can't be set")
	vp := reflect.ValueOf(&fVar)
	fmt.Printf("is float canSet: %v canAddr %v\n", vp.Elem().CanSet(), vp.Elem().CanAddr())
	vp.Elem().SetFloat(2.33333)

	println(fVar)
}
