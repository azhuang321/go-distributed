# 数组概述

### 一 数组概念

> 数组（array）：有限个相同类型元素组成的有序集合

数组是最简单的数据结构，基本所有的语言都提供了原生的数组数据结构。数组的长度是固定的，且存储的元素类型也都是相同，索引从0开始。其存储结构如下所示：  


![](https://github.com/overnote/over-algorithm/raw/master/images/structure/array-01.svg)

数组本质上是在物理上一组连续的内存上存储的数据（顺序存储），如果要移动一个元素，其相关元素也都需要一一移动，如图所示：

![](https://github.com/overnote/over-algorithm/raw/master/images/structure/array-02.svg)

在早期的语言中，数组并不支持在运行期间改变大小，必须预定义数组的容量，比如C语言，一些近代语言如Java,JS是支持数组的动态定义的，即长度可变，所以数组可以分为：

* 静态数组：编译时确定数组的长度，为了防止空间不足，所以尽量将数组的长度定义的大点，但是容易造成内存浪费
* 动态数组：不需要在编译时确定长度，而是在运行过程中确定。

### 二 随机读取

> 随机读取：使用下标读取元素的方式

具备随机读取特性的数据操作，其时间复杂度为O\(1\)，如数组的元素获取、更新。

### 三 数组的常见操作

数组的常见操作：

* 获取元素：`arr[index]`，时间复杂度为O\(1\)
* 更新元素：`arr[index] = newElem`，时间复杂度为O\(1\)
* 元素的插入与删除：由于数组元素紧紧相邻，插入/删除元素需要移动其周边元素一位，时间复杂度为O\(n\)。不同编程语言为数组提供了不同的插入删除操作API。

### 四 总结

* 数组的优势：具备随机读取特性，可以在常量时间内访问元素
* 数组的劣势：插入、删除操作会造成大量元素被迫移动

所以：数组适合读操作多、写操作少的场景。

### 五 Go中的数组

#### 5.1 数组的声明

同大多语言的数组一样，Go的数组长度定义后不可更改，长度使用 len\(\) 获取。

```go
var arr1 [10]int					//定义长度为10的整型数组，很少这样使用
arr2 [5]int := [5]int{1,2,3,4,5}	//定义并初始化
arr3 := [5]int{1,2,3,4,5}			//自动推导并初始化
arr4 := [5]int{1,2}					//指定总长度，前几位被初始化，没有的使用零值
arr5 := [5]int{2:10, 4:11}			//有选择的初始化，没被初始化的使用零值
arr6 := [...]int{2,3,4}				//自动计算长度
```

#### 5.2 数组元素获取

```go
arr[:]      代表所有元素
arr[:5]     代表前五个元素，即区间的左闭右开
arr[5:]     代表从第5个开始（不包含第5个）
len(arr)    数组的长度
```

贴士：上述操作会引发类型的变化，数组将会转化为Go中新的数据类型切片

#### 5.3 数组的遍历

方式一：for循环遍历

```go
arr := [3]int{1,2,3}

for i := 0; i < len(arr); i++ {
	fmt.Println(arr[i])
}
```

方式二：for-range遍历

```go
arr := [3]int{1,2,3}

for k, v := range arr {
	fmt.Println(k)	//元素位置	
	fmt.Println(v)	//元素值
}
```

#### 5.4 数组使用注意事项

数组创建完长度就固定，不可以再追加元素；长度是数组类型的一部分，因此`[3]int`与`[4]int`是不同的类型；数组之间的赋值是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该函数的副本，而不是他的指针。

