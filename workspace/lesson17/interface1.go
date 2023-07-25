package main

import "fmt"

// all animals can speak
type Animal interface {
	speak()
	eat()
}

// cat
type Cat struct {
	name string
	age  int
}

func (cat Cat) speak() {
	fmt.Println("cat miaomiaomiao")
}

func (cat Cat) eat() {
	fmt.Println("cat eating")
	cat.age++
}

// dog
type Dog struct {
	name string
	age  int
}

func (dog *Dog) speak() {
	fmt.Println("dog wangwangwang")
}

// 不会改变结构体变量的值
func (dog *Dog) eat() {
	fmt.Println("dog eating")
	dog.age--
}

func main() {
	/*
	   Cat实现speak方法用的是值接受者，给animal赋值的时候
	   使用值或者指针都可以，var animal Animal = &Cat{"gaffe", 1}
	*/
	cat := Cat{
		name: "gaffe",
		age:  1,
	}
	var animal Animal = &cat
	animal.speak() // cat miaomiaomiao
	animal.eat()
	fmt.Println("cat's age is:", cat.age) //即便赋值的是指针，也不会改变实际的成员变量值

	/*
	   因为Dog的speak方法用的是指针接受者，因此给interface赋值的时候，要赋指针
	*/
	dog := Dog{
		name: "caiquan",
		age:  3,
	}
	animal = &dog
	animal.speak() // dog wangwangwang
	animal.eat()
	fmt.Println("dog's age is:", dog.age) //这里会改变成员变量值
}
