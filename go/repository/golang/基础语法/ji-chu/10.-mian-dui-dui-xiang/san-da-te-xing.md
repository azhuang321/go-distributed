# 三大特性

### 一 面向对象三大特性

**1.1 封装**

封装：把抽象出的字段和对字段的操作封装在一起,数据被保护在内部,程序的其它包只有通过被授权的操作\(方法\),才能对字段进行修改，其作用有：

* 隐藏实现细节
* 可以对数据进行验证，保证安全合理

Golang对面向对象做了极大简化，并不强调封装特性，下列示例进行模拟实现：

在`person`包下新建`person.go`文件：

```go
package person

import "fmt"

type person struct {
	Name string
	age int			//年龄是隐私，不允许其他包访问
}

//工厂函数（类似构造函数）
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {		//校验
		p.age = age
	} else {
		fmt.Println("年龄不合法")
	}
}

func (p *person) GetAge() int {
	return p.age
}
```

`main.go`文件操作person：

```go
package main

import (
	"demo/person"					// demo是go mod模式下，整体项目名
	"fmt"
)

func main() {
	p := person.NewPerson("Tom")
	p.SetAge(18)
	fmt.Println(p)
}
```

**1.2 继承**

在 Golang 中，如果一个 struct 嵌套了另一个匿名结构体，那么这个结构体可以直接访 问匿名结构体的字段和方法，从而实现了继承特性。

```go
package main

import (
	"fmt"
)

type Father struct {
	Name string
	age int
}
func (f *Father) run() {
	fmt.Println(f.Name + " like running...")
}

type Son struct {
	Father              //嵌套匿名结构体
}

func main() {

	var s Son

	//s.Father.Name = "Tom"
	//s.Father.age = 10     		//可以访问未导出属性
	//s.Father.run()          	//可以访问未导出方法

	//上述可以简写为：
	s.Name = "Tom"
	s.age = 10
	s.run()

}
```

注意：

* 当结构体和匿名结构体有相同的字段或者方法时，**编译器采用就近访问原则访问**，如果希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分。
* 结构体嵌入多个匿名结构体，如果两个匿名结构体有相同的字段和方法\(同时结构体本身没有同名的字段和方法\)，访问时必须明确指定匿名结构体名字，否则编译报错。
* 如果一个 struct 嵌套了一个有名结构体，这种模式就是组合，如果是组合关系，那么在访问组合的结构体的字段或方法时，必须带上结构体的名字。

关于多重继承：如果一个 struct 嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方法，从而实现多重继承。

```go
package main

import (
	"fmt"
)

type Father1 struct {
	Name string
	age int
}
func (f *Father1) run() {
	fmt.Println(f.Name + " like running...")
}

type Father2 struct {
	Like string
}


type Son1 struct {
	Father1
	Father2
}

type Son2 struct {
	*Father1
	*Father2
}

func main() {

	s1 := &Son1 {
		Father1{
			Name: "Tom",
			age: 10,
		},
		Father2{
			Like: "伏特加",
		},
	}

	fmt.Println(s1)

	s2 := &Son2{
		&Father1{
			Name: "Tom",
			age: 10,
		},
		&Father2{
		 	Like: "伏特加",
		},
	}
	fmt.Println(s2.Father1)

}
```

输出结果：

```go
&{{Tom 10} {伏特加}}
&{Tom 10}
```

**1.3 多态**

多态与接口（interface）有关联，参见接口章节

