# 构造函数与方法

### 一 面向对象初识

**1.1 模拟构造函数**

Go和传统的面向对象语言如Java有着很大区别。结构体没有构造函数初始化功能，可以通过以下方式模拟：

```go
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func NewPersonByName(name string) *Person {
	return &Person{
		Name: name,
	}
}

func NewPersonByAge(age int) *Person {
	return &Person{
		Age: age,
	}
}

func main() {

	p := NewPersonByName("zs")
	fmt.Println(p)						// {zs 0}

}
```

贴士：因为Go没有函数重载，为了避免函数名字冲突，使用了`NewPersonByName`和`NewPersonByAge`两个不同的函数表示不同的`Person`构造过程。

**1.2 父子关系结构体初始化**

Person可以看做父类，Student是子类，子类需要继承父类的成员：

```go
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

type Student struct {
	Person
	ClassName string
}

//构造父类
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age: age,
	}
}

//构造子类
func NewStudent(classname string) *Student {
	p := &Student{}
	p.ClassName = classname
	return p
}


func main() {

	s := NewStudent("一班")
	fmt.Println(s)						// &{{ 0} 一班}

}
```

**1.3 Go中的面向对象初识**

在Go中，可以给任意类型（除了指针）添加相应方法：

```go
type Interger int

func (i Interger) Less (j Interger) bool {
	return i < j
}

func main() {
	var i Interger = 1
	fmt.Print(i.Less(5))
}
```

### 二 方法

**2.1 方法**

Golang 中的方法是作用在指定的数据类型上的\(即:和指定的数据类型绑定\)，因此自定义类型，都可以有方法，而不仅仅是 struct。

方法的声明和调用：

```go
func (recevier type) methodName(参数列表) (返回值列表){ 
    //方法体
    return 返回值
}
```

方法与函数的示例：

```go
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

// 一个run函数
func run(p *Person, name string) {
	p.Name = name
	fmt.Println("函数 run...", p.Name)
}

// 一个run方法
func (p *Person)run() {
	fmt.Println("方法 run...", p.Name)
}

func main() {

	// 实例化一个对象（结构体）
	p1 := &Person{
		"ruyue",
		10,
	}

	// 执行一个普通方法
	run(p1, "张三")			// 输出 函数 run... 张三

	// 执行方法
	p1.run()						// 输出 方法 run... 张三

}
```

**2.2 Go方法本质**

Go的方法是一种作用于特定类型变量的函数，这种特定类型的变量叫做接收器（Receiver）。如果特定类型理解为结构体或者“类”时，接收器就类似于其他语言的this或者self。

在Go中，接收器可以是任何类型，不仅仅是结构体，依此我们看出，Go中的方法和其他语言的方法类似，但是Go语言的接收器强调方法的作用对象是实例。

方法与函数的区别就是：函数没有作用对象。

**指针接收器传入的是 struct 本身**，指针接收器可以读写 struct 中的内容，在方法结束后，修改都是有效的。

**非指针接收器传入的是 struct 的 copy 副本**，非指针接收器只能读取 struct 中的数据但是不能写入，如果写入的话也只是写入到 struct 的备份中而已。

示例如下:

```go
package main

import "fmt"

type student struct {
	age int8
}

//指针接收器
func(s *student) ageAdd1() {
	s.age += 1
}

//非指针接收器
func(s student) ageAdd2() {
	s.age += 1
}

func main() {
	student := new(student)

	student.ageAdd1()
	fmt.Println(student.age) // 1 传入指针，原值 + 1，为 1

	student.ageAdd1()
	fmt.Println(student.age) // 2 传入指针，原值 + 1，为 2

	student.ageAdd2()
	fmt.Println(student.age) // 2 传入复制体，复制体 + 1，所以原值还是 2
}
```

一般情况下，小对象由于复制时速度较快，适合使用非指针接收器，大对象因为复制性能较低，适合使用指针接收器，此时再接收器和参数之间传递时不进行复制，只传递指针。

