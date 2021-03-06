# 指针

### 一 指针

#### 1.1 指针的创建

Go保留了指针，代表某个内存地址，默认值为 `nil` ，使用 `&` 取变量地址，通过 `*` 访问目标对象。

简单示例：

```text
	var a int = 10
	fmt.Println("&a=", &a)			// 0xc000096008 一个十六进制数

	var p *int = &a
	fmt.Println("*p=", *p)			// 10
```

注意：

* Go同样支持多级指针，如 `**T`
* 空指针：声明但未初始化的指针
* 野指针：引用了无效地址的指针，如：`var p *int = 0`，`var p *int = 0xff00`\(超出范围\)
* Go中直接使用`.`访问目标成员

#### 1.2 指针使用示例：实现变量值交换

```text
	func swap(p1,p2 *int) {
	*p1,*p2 = *p2,*p1
	}
```

#### 1.3 结构体指针

示例：

```text
	type User struct{
		name string
		age int
	}

	func main() {
		var u = User{
			name:"lisi",
			age: 18,
		}
		p := &u
		fmt.Println(u.name)		//输出李四
		fmt.Println(p.name)		//输出李四
	}
```

#### 1.4 Go不支持指针运算

由于垃圾回收机制的存在，指针运算造成许多困扰，所以Go直接禁止了指针运算

```text
	a := 1
	p := &a
	p++        //报错：non-numeric type *int
```

#### 1.5 new\(\)函数使用

new\(\)函数可以在 heap堆 区申请一片内存地址空间：

```text
	var p *bool
	p = new(bool)
	fmt.Println(*p)		// false
```

