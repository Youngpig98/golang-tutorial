package main

import (
	"fmt"
	"reflect"
	"time"
)

func makeTimeFunc(f interface{}) interface{} {
	tf := reflect.TypeOf(f)
	vf := reflect.ValueOf(f)

	if tf.Kind() != reflect.Func {
		panic("expect a function")
	}
	wrapper := reflect.MakeFunc(tf, func(args []reflect.Value) (result []reflect.Value) {
		start := time.Now()
		result = vf.Call(args)
		end := time.Now()
		fmt.Printf("The function takes %v\n", end.Sub(start))
		return result
	})
	return wrapper.Interface() //相当于valueof的反向操作
}

func TimeMe(str string) (i int64) {
	fmt.Println("Doing something", str)
	time.Sleep(4 * time.Second)
	fmt.Println("finished")
	return 32
}

func main() {
	timedFunc := makeTimeFunc(TimeMe).(func(string) int64)
	i := timedFunc("hello")
	fmt.Println(i)
}
