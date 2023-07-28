# Go反射机制

## 一、理解变量的内在机制

1. 类型信息，元信息，是预先定义好的，静态的。

2. 值信息，程序进行过程中，动态变化的。

## 二、反射和空接口

1. 空接口相当于一个容器，能接受任何东西。

2. 那怎么判断空接口变量存储的是什么类型呢？之前有使用过类型断言，这只是一个比较基础的方法

3. 如果想获取存储变量的类型信息和值信息就要使用反射机制，所以反射是什么？ 反射就是动态的获取变量类型信息和值信息的机制。

## 三、怎么利用反射分析空接口里面的信息呢？

1. 首先利用的是GO语言里面的Reflect包

2. 利用包里的TypeOf方法可以获取变量的类型信息

```go
func printMeta(obj any) {
	// pair: <value, type>
	t := reflect.TypeOf(obj)
	name := t.Name()
	kind := t.Kind()
	value := reflect.ValueOf(obj)
	fmt.Printf("Type: %s Type.Name: %s Kind: %s Value: %v\n", t, name, kind, value)
}
```

​	利用Kind() 可以获取t的类型，如代码所示，这里可以判断a是Int64还是string, 像下面一样使用：

```go
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
}func` `main() {``  ``var` `x ``int64` `= ``3``  ``reflect_example(x)`` ` `  ``var` `y ``string` `= ``"hello"``  ``reflect_example(y)``}
```

3. 利用包里的ValueOf方法可以获取变量的值信息

```go
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
```

​	利用ValueOf方法可以得到变量的值信息，ValueOf返回的是一个Value结构体类型，有趣的是，可以使用 v.Type() 获取该变量的类型，和上面reflect.TypeOf() 获取的结果一样。

​	这里存在一个问题，如果传进去一个类型，使用了错误的解析，那么将会在运行的时候报错， 例如将 一个string类型强行的v.Int()。

​	既然值类型是动态的，能取到保存的值，同样可以设置值。在反射里面有很多set的方法，例如SetFloat、SetInt()、SetString()等可以帮助我们设置值。

​	下面的例子，我想把 x设置为 6.28，但是会报错！

```go
func main() {
    var x float64 = 3.14
    v := reflect.ValueOf(x)
    v.SetFloat(6.28) 
    fmt.Printf("After Set Value is %f", x)
}
```

错误结果：

> panic: reflect: reflect.Value.SetFloat using unaddressable value
> ......

​	结果上说明是不可设置的，为什么呢？ 因为我们的x是一个值类型，而值类型的传递是拷贝了一个副本，当 v := reflect.ValueOf(x) 函数通过传递一个 x 拷贝创建了 v，那么 v 的改变并不能更改原始的 x。要想 v 的更改能作用到 x，那就必须传递 x 的地址 `v = reflect.ValueOf(&x)`。修改程序如下：

```go
func main() {
    var x float64 = 3.14
    v := reflect.ValueOf(&x)
    v.SetFloat(6.28)
    fmt.Printf("After Set Value is %f", x)
}
```

​	结果：依然报错！为什么传了地址还报错？因为&x是地址了，所以它的类型就变了，可以通过v.Type()，看下改变成了什么：

```go
func main() {
    var x float64 = 3.14
    v := reflect.ValueOf(&x)
    fmt.Printf("type of v is %v", v.Type())   //打印的结果是：type of v is *float64
}
```

​	由程序可以知道，这个返回的是一个指针类型的。所以上面SetFloat才会失败，那怎么做？

​	我们正常的赋值，如果是地址的话，例如下面：一般我们都会对*y进行赋值, *的意思就是往这个地址里面赋值。

```go
var y *float64 = new(float64)
*y = 10.12
fmt.Printf("y = %v", *y)
```

​	同样的，我们在反射里面也可以取地址，需要通过 Elem() 方法进行取地址。再次修改程序

```go
func main() {
    var x float64 = 3.14
    v := reflect.ValueOf(&x)
    fmt.Printf("type of v is %v\n", v.Type())
    v.Elem().SetFloat(6.28)
    fmt.Printf("After set x is %v", x)
}
```

结果为：

> type of v is *float64
> After set x is 6.28

## 四、利用反射获取结构体里面的方法和调用。

1. 获取结构体的字段

​	我们可以通过上面的方法判断一个变量是不是结构体。

​	可以通过 NumField() 获取所有结构体字段的数目、进而遍历，通过Field()方法获取字段的信息。

```go
type Student struct {
    Name  string
    Sex   int
    Age   int
    Score float32
}
 
func main() {
    //创建一个结构体变量
    s := Student{
        Name:  "BigOrange",
        Sex:   1,
        Age:   10,
        Score: 80.1,
    }
 
    v := reflect.ValueOf(s)
    t := v.Type()
    kind := t.Kind()
     
    //分析s变量的类型，如果是结构体类型，那么遍历所有的字段
    switch kind {
    case reflect.Int64:
        fmt.Printf("s is int64\n")
    case reflect.Float32:
        fmt.Printf("s is int64\n")
    case reflect.Struct:
        fmt.Printf("s is struct\n")
        fmt.Printf("field num of s is %d\n", v.NumField())
        //NumFiled()获取字段数，v.Field(i)可以取得下标位置的字段信息，返回的是一个Value类型的值
        for i := 0; i < v.NumField(); i++ {
            field := v.Field(i)
            //打印字段的名称、类型以及值
            fmt.Printf("name:%s type:%v value:%v\n",
                t.Field(i).Name, field.Type().Kind(), field.Interface())
        }
    default:
        fmt.Printf("default\n")
    }
}
```

​	执行结果：

> s is struct
> field num of s is 4
> name:Name type:string value:BigOrange
> name:Sex type:int value:1
> name:Age type:int value:10
> name:Score type:float32 value:80.1

​	这里需要说明几个问题：

​		(1) 打印字段名称的时候，使用的是t.Field(i).Name ，Name是静态的，所以属于类型的信息

​		(2) 打印值的时候，这里将field.Interface()实际上相当于ValueOf的反操作（可以参考这篇文章https://www.jb51.net/article/255856.htm），所以才能把值打印出来

​		(3) 此外如果Student中的Name字段变为name（私有），那么则会报错，不能反射出私有变量 错误信息 “panic: reflect.Value.Interface: cannot return value obtained from unexported field or method”

2. 对结构体内的字段进行赋值操作

​	参考下面的代码，对上面的Student进行赋值操作：

```go
func main() {
    s := Student{
        Name:  "BigOrange",
        Sex:   1,
        Age:   10,
        Score: 80.1,
    }
 
    fmt.Printf("Name:%v, Sex:%v,Age:%v,Score:%v \n", s.Name, s.Sex, s.Age, s.Score)
    v := reflect.ValueOf(&s)  //这里传的是地址！！！
 
    v.Elem().Field(0).SetString("ChangeName")
    v.Elem().FieldByName("Score").SetFloat(99.9)
 
    fmt.Printf("Name:%v, Sex:%v,Age:%v,Score:%v \n", s.Name, s.Sex, s.Age, s.Score)
}
```

结果：

> Name:BigOrange, Sex:1,Age:10,Score:80.1
> Name:ChangeName, Sex:1,Age:10,Score:99.9

3. 获取结构体里面的方法

​	可以通过NumMethod()获得接头体里面的方法数量，然后遍历通过Method()获取方法的具体信息。如下代码所示：

```go
//新增-设置名称方法
func (s *Student) SetName(name string) {
     fmt.Printf("有参数方法 通过反射进行调用:%v\n", s)
     s.Name = name
}
//新增-打印信息方法
func (s *Student) PrintStudent() {
    fmt.Printf("无参数方法 通过反射进行调用:%v\n", s)
}
 
func main() {
    s := Student{
        Name:  "BigOrange",
        Sex:   1,
        Age:   10,
        Score: 80.1,
    }
 
    v := reflect.ValueOf(&s)
    //取得Type信息
    t := v.Type()
     
    fmt.Printf("struct student have %d methods\n", t.NumMethod())
 
    for i := 0; i < t.NumMethod(); i++ {
        method := t.Method(i)
        fmt.Printf("struct %d method, name:%s type:%v\n", i, method.Name, method.Type)
    }
}
```

输出：

> struct student have 2 methods
> struct 0 method, name:PrintStudent type:func(*main.Student)
> struct 1 method, name:SetName type:func(*main.Student, string)

​	从结果中看到我们可以获取方法的名称以及签名信息，并且这个方法的输出顺序是按照字母排列的。

​	并且输出结果可以看到一个有趣的现象：结构体的方法其实也是通过函数实现的例如 func(s *Student) SetName(name string) 这个方法，反射之后的结果就是 func(*main.Student , string) 实际上把Student当参数了。

​	此外还可以通过反射来调用这些方法。想要通过反射调用结构体里面的方法，首先要知道方法调用时一个动态的，所以要先通过ValueOf获取值，然后通过获取的值进行方法的调用 ，通过 value里面的Method方法 返回一个方法，然后通过Call方法调用，Call是参数是一个切片，也就是参数的列表。如下代码展示了如何调用有参数的方法和无参数的方法：

```go
func main() {
    s := Student{
        Name:  "BigOrange",
        Sex:   1,
        Age:   10,
        Score: 80.1,
    }
 
    v := reflect.ValueOf(&s)  //传指针
 
    //通过reflect.Value获取对应的方法并调用
    m1 := v.MethodByName("PrintStudent")
    var args []reflect.Value
    m1.Call(args)
 
    m2 := v.MethodByName("SetName")
    var args2 []reflect.Value
    name := "stu01"
    nameVal := reflect.ValueOf(name)
    args2 = append(args2, nameVal)
    m2.Call(args2)
    m1.Call(args)
}
```

执行结果：

> 无参数方法 通过反射进行调用:&main.Student{Name:"BigOrange", Sex:1, Age:10, Score:80.1}
> 有参数方法 通过反射进行调用:&main.Student{Name:"BigOrange", Sex:1, Age:10, Score:80.1}
> 无参数方法 通过反射进行调用:&main.Student{Name:"stu01", Sex:1, Age:10, Score:80.1}

上面格式打印：

- %v 相应值的默认格式。 Printf("%v", people) {zhangsan}，
- %+v 打印结构体时，会添加字段名 Printf("%+v", people) {Name:zhangsan}
- %#v 相应值的Go语法表示 Printf("#v", people) main.Human{Name:"zhangsan"}

## 五、怎么获取结构体里tag的信息。

​	有时候我们在类型上面定义一些tag，例如使用json和数据库的时候。Field()方法返回的StructField结构体中保存着Tag信息，并且Tag信息可以通过一个Get(Key)的方法获取出来，如下代码所示：

```go
type Student struct {
    Name string `json:"jsName" db:"dbName"`
}
 
func main() {
    s := Student{
        Name: "BigOrange",
    }
    v := reflect.ValueOf(&s)
    t := v.Type()
    field0 := t.Elem().Field(0)
    fmt.Printf("tag json=%s\n", field0.Tag.Get("json"))
    fmt.Printf("tag db=%s\n", field0.Tag.Get("db"))
}
```

结果：

> tag json=jsName
> tag db=dbName

## 六、应用场景

1. 序列化和反序列化，比如json, protobuf等各种数据协议

2. 各种数据库的ORM，比如gorm，sqlx等数据库中间件

3. 配置文件解析相关的库，比如yaml、ini等



​	实际上，Go语言的反射机制无法直接获取变量的名称，因为在运行时，变量名已经被编译器丢弃了。反射只能获取变量的类型、值和一些其他相关信息，而无法获取变量的名称。如果您想要获取变量的名称，需要通过手动指定变量名的方式来实现。