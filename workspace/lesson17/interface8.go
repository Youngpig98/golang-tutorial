package main

import "fmt"

// 定义一个接口
type Speaker interface {
	Speak()
}

// 定义一个实现了 Speaker 接口的结构体
type Dog struct {
}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

// 定义一个结构体，包含接口类型的字段
type AnimalShelter struct {
	Pet Speaker
}

func main() {
	// 创建一个 AnimalShelter 实例，并设置 Pet 字段为 Dog 类型的实例
	shelter := AnimalShelter{
		Pet: Dog{},
	}

	// 调用 Pet 字段的 Speak 方法
	shelter.Pet.Speak() // 输出: Woof!
}