# 接口interface

* 定义：接口是一种抽象的类型，是一组method的集合，里头只有method方法，没有数据成员。当两个或两个以上的类型都有相同的处理方法时才需要用到接口。先定义接口，然后多个struct类型去实现接口里的方法，就可以通过接口变量去调用struct类型里实现的方法。

  比如动物都会叫唤，那可以先定义一个名为动物的接口，接口里有叫唤方法speak，然后猫和狗这2个struct类型去实现各自的speak方法。

* 语法：

  ```go
  // 定义接口
  type interface_name interface {
    method_name1([参数列表]) [返回值列表]
    method_name2([参数列表]) [返回值列表]
    method_nameN([参数列表]) [返回值列表]
  }
  
  // 定义结构体类型
  type struct_name struct {
      data_member1 data_type
      data_member2 data_type
      data_memberN data_type
  }
  
  // 实现接口interface_name里的方法method_name1
  func(struct_var struct_name) method_name1([参数列表])[返回值列表] {
      /*具体方法实现*/
  }
  
  // 实现接口interface_name里的方法method_name2
  func(struct_var struct_name) method_name2([参数列表])[返回值列表] {
      /*具体方法实现*/
  }
  
  /* 实现接口interface_name里的方法method_name3
  注意：下面用了指针接受者。函数可以使用值接受者或者指针接受者，上面的method_name1和method_name1使用的是值接受者。
  如果用了指针接受者，那给interface变量赋值的时候要传指针
  */
  func(struct_var *struct_name) method_name3([参数列表])[返回值列表] {
      /*具体方法实现*/
  }
  
  ```

  

* 示例：

  ```go
  package main
  
  import "fmt"
  
  // all animals can speak
  type Animal interface {
      speak()
  }
  
  // cat
  type Cat struct {
      name string
      age int
  }
  
  func(cat Cat) speak() {
      fmt.Println("cat miaomiaomiao")
  }
  
  // dog
  type Dog struct {
      name string
      age int
  }
  
  func(dog *Dog) speak() {
      fmt.Println("dog wangwangwang")
  }
  
  
  func main() {
      var animal Animal = Cat{"gaffe", 1}
      animal.speak() // cat miaomiaomiao
      
      /*
      因为Dog的speak方法用的是指针接受者，因此给interface赋值的时候，要赋指针
      */
      animal = &Dog{"caiquan", 2}
      animal.speak() // dog wangwangwang
  }
  ```

* struct结构体类型在实现interface里的所有方法时，关于interface变量赋值有2个点要**注意**

  * 只要有某个方法的实现使用了**指针接受者**，那给包含了这个方法的interface变量赋值的时候要**使用指针**。比如上面的Dog类型要赋值给Animal，必须使用指针，因为Dog实现speak方法用了指针接受者。

  * 如果全部方法都使用的是值接受者，那给interface变量赋值的时候用值或者指针都可以。比如上面的例子，animal的初始化用下面的方式一样可以：

    ```go
    var animal Animal = &Cat{"gaffe", 1}
    ```

* **多个struct类型可以实现同一个interface**：多个类型都有共同的方法(行为)。比如上面示例里的猫和狗都会叫唤，猫和狗就是2个类型，叫唤就是speak方法。

* **一个struct类型可以实现多个interface**。比如猫这个类型，既是猫科动物，也是哺乳动物。猫科动物可以是一个interface，哺乳动物可以是另一个interface，猫这个struct类型可以实现猫科动物和哺乳动物这2个interface里的方法。

  ```go
  package main
  
  import "fmt"
  
  
  // interface1，猫科动物的共同行为
  type Felines interface {
      feet() 
  }
  
  // interface2, 哺乳动物的共同行为
  type Mammal interface {
      born()
  }
  
  // 猫既是猫科动物也是哺乳动物，2个行为都实现
  type Cat struct {
      name string
      age int
  }
  
  func(cat Cat) feet() {
      fmt.Println("cat feet")
  }
  
  func(cat *Cat) born() {
      fmt.Println("cat born")
  }
  
  func main() {
      cat := Cat{"rich", 1}
      var a Felines = cat
      a.feet()
      
      var b Mammal = &cat
      b.born()
  }
  ```

  

* interface可以嵌套：一个interface里包含其它interface

  ```go
  package main
  
  import "fmt"
  
  
  // interface1
  type Felines interface {
      feet() 
  }
  
  // interface2, 嵌套了interface1
  type Mammal interface {
      Felines
      born()
  }
  
  // 猫实现Mammal这个interface里的所有方法
  type Cat struct {
      name string
      age int
  }
  
  func(cat Cat) feet() {
      fmt.Println("cat feet")
  }
  
  func(cat *Cat) born() {
      fmt.Println("cat born")
  }
  
  func main() {
      cat := Cat{"rich", 1}
      /*Mammal有feet和born方法，2个都可以调用*/
      var a Mammal = &cat
      a.feet()
      a.born()
      
      var b Felines = cat
      b.feet()
      // b.born() 调用这个会编译报错，因为Felines没有born方法
  }
  ```

  

* 空接口interface                      **在Go1.18中引入了any关键字，它其实就是空接口的一个别名**

  * 如果空interface作为函数参数，可以接受任何类型的实参

    * 语法

      ```go
      func function_name(x interface{}) {
          do sth
      }
      ```

      

    * 示例

      ```go
      package main
      
      import "fmt"
      
      
      type Cat struct {
          name string
          age int
      }
      
      // 打印空interface的类型和具体的值
      func print(x interface{}) {
          fmt.Printf("type:%T, value:%v\n", x, x)
      }
      
      func main() {
          // 传map实参给空接口
          dict := map[string]int{"a":1}
          print(dict) // type:map[string]int, value:map[a:1]
          
          // 传struct实参给空接口
          cat := Cat{"nimo", 2}
          print(cat) // type:main.Cat, value:{nimo 2}
      }
      ```

      

  * 如果空interface作为变量，可以把任何类型的变量赋值给空interface

    * 语法 

      ```go
      var x interface{} // 空接口x
      ```

    * 示例

      ```go
      package main
      
      import "fmt"
      
      
      type Cat struct {
          name string
          age int
      }
      
      // 打印空interface的类型和具体的值
      func print(x interface{}) {
          fmt.Printf("type:%T, value:%v\n", x, x)
      }
      
      func main() {
          // 定义空接口x
          var x interface{}
          // 将map变量赋值给空接口x
          x = map[string]int{"a":1}
          print(x) // type:map[string]int, value:map[a:1]
          
          // 传struct变量估值给空接口x
          cat := Cat{"nimo", 2}
          x = cat
          print(x) // type:main.Cat, value:{nimo 2}
      }
      ```
    
  * 空接口作为map的值，可以实现map的value是不同的数据类型

    * 语法

      ```go
      // 定义一个map类型的变量，key是string类型，value是空接口类型
      dict := make(map[string]interface{}) 
      ```

      

    * 示例

      ```go
      package main
      
      import "fmt"
      
      
      func main() {
          // 定义一个map类型的变量，key是string类型，value是空接口类型
          dict := make(map[string]interface{})
          // value可以是int类型
          dict["a"] = 1 
          // value可以是字符串类型
          dict["b"] = "b"
          // value可以是bool类型
          dict["c"] = true
          fmt.Println(dict) // map[a:1 b:b c:true]
          fmt.Printf("type:%T, value:%v\n", dict["b"], dict["b"]) // type:string, value:b
      }
      ```

      

  * x.(T)

    * 断言：断言接口变量x是T类型

      * 语法：value是将x转化为T类型后的变量，ok是布尔值，true表示断言成功，false表示断言失败

        ```go
        // x是接口变量，如果要判断x是不是
        value, ok := x.(string)
        ```

      * 示例

        ```go
        var x interface{}
        x = "a"
        // 断言接口变量x的类型是string
        v, ok := x.(string)
        if ok {
            // 断言成功
            fmt.Println("assert true, value:", v)
        } else{
            // 断言失败
        	fmt.Println("assert false")
        }
        ```
      
    * 动态判断数据类型

      ```go
      package main
      
      import "fmt"
      
      func checkType(x interface{}) {
          /*动态判断x的数据类型*/
          switch v := x.(type) {
          case int:
              fmt.Printf("type: int, value: %v\n", v)
          case string:
              fmt.Printf("type: string，value: %v\n", v)
          case bool:
              fmt.Printf("type: bool, value: %v\n", v)
          case Cat:
              fmt.Printf("type: Cat, value: %v\n", v)
          case map[string]int:
              fmt.Printf("type: map[string]int, value: %v\n", v)
              v["a"] = 10
          default:
              fmt.Printf("type: %T, value: %v\n", x, x)
          }
      }
      
      type Cat struct {
          name string
          age int
      }
      
      func main() {   
          var x interface{}
          x = "a"
          checkType(x) //type: string，value: a
          
          x = Cat{"hugo", 3}
          checkType(x) // type: Cat, value: {hugo 3}
      
          /*在checkType里对map做修改
          会影响外面的实参x
          */
          x = map[string]int{"a":1}
          checkType(x) // type: map[string]int, value: map[a:1]
          fmt.Println(x) // map[a:10]
      }
      ```
      

## 注意事项

* **如果把一个结构体变量赋值给interface变量，那结构体需要实现interface里的所有方法，否则会编译报错**：xx does not implement yy，表示结构体xx没有实现接口yy如下所示，在接口中包含WriteData方法和CanWrite方法，只有在file结构体绑定方法WriteData和CanWrite之后，才能正常使用接口。
* go语言的接口实现是隐式的，我们无须让实现接口的（结构体）类型知道有哪些接口（需要被实现），这个称为非侵入式设计。我们在编写（接口）方法时，无法预测未来有哪些（接口）方法会变成接口。一旦某个接口被创建出来，要求旧代码来实现这个接口，就需要修改旧代码的派生部分，这一般会造成雪崩式的重新编译。
* 传统的派生接口和类构建模式，使类型之间有强耦合的父子关系。这种关系一般以类派生图的方式运行。随着系统功能的不断增加，这棵“派生树“会越来越复杂。非侵入式设计让go程序实现的所有类型都是平行的、组合的。因此，使用go时，不需要也不可能有类派生，开发者唯一需要关注的是“我需要什么，我能实现什么”。



## Go 接⼝口最佳实践

- 倾向于使⽤用⼩小的接⼝口定义，很多接⼝只包含一个方法

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
```

- 较大的接口定义，可以由多个小接口定义组合而成

```go
type ReadWriter interface {
	Reader
	Writer
}
```

- 只依赖于必要功能的最小接口

```go
func StoreData(reader Reader) error {
…
}
```



参考：[Methods that satisfy interfaces in golang](https://suraj.io/post/golang-methods-interfaces/)
