# 切片

### 一 切片创建

切片\(slice\)解决了数组长度不能扩展，以及基本类型数组传递时产生副本的问题。

常用创建方式：

```text
var s1 []int				// 和声明数组一样，只是没有长度，但是这样做没有意义，因为底层的数组指针为nil
s2 := []byte {'a','b','c'}
fmt.Println(s1)				//输出 []
fmt.Print(s2)				//输出 [97 98 99]
```

使用make函数创建：

```text
slice1 := make([]int,5)		// 创建长度为5，容量为5，初始值为0的切片
slice2 := make([]int,5,7)	// 创建长度为5，容量为7，初始值为0的切片
slice3 := []int{1,2,3,4,5}	// 创建长度为5，容量为5，并已经初始化的切片
```

从数组创建：slice可以从一个数组再次声明。slice通过array\[i:j\]来获取，其中i是数组的开始位置，j是结束位置，但不包含array\[j\]，它的长度是j-i:

```text
// 声明一个含有10个元素元素类型为byte的数组
var arr = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

// 声明两个含有byte的slice
var a, b []byte

// a指向数组的第3个元素开始，并到第五个元素结束，现在a含有的元素: ar[2]、ar[3]和ar[4]
a = arr[2:5]		
// b是数组arr的另一个slicre,b的元素是：ar[3]和ar[4]
b = arr[3:5]			
```

注意：声明数组时，方括号内写明了数组的长度或使用...自动计算长度，而声明slice时，方括号内没有任何字符。

从切片创建：

```text
oldSlice := []int{1,2,3}
newSlice := oldSlice[:6]	//基于切片前6个元素创建，没有的默认0
```

注意：如果选择的旧切片长度超出了旧切片的cap\(\)值（切片存储长度），则不合法。

### 二 切片常见操作

**2.1 切片常见内置函数**

切片常用内置函数：

```text
len()			返回切片长度
cap()			返回切片底层数组容量
append()		对切片追加元素
func copy(dst, src []Type) int
				将src中数据拷贝到dst中，返回拷贝的元素个数
```

切片空间与元素个数：

```text
slice1 := make([]int, 5, 10)
fmt.Println(len(slice1))			// 5
fmt.Println(cap(slice1))			// 10
fmt.Println(slice1)					// [0 0 0 0 0]
```

切片操作

```text
//切片增加
slice1 = append(slice1,1,2)
fmt.Println(slice1)						//输出[0 0 0 0 0 1 2]

//切片增加一个新切片
sliceTemp := make([]int,3)
slice1 = append(slice1,sliceTemp...)
fmt.Println(slice1)						//输出[0 0 0 0 0 1 2 0 0 0]

//切片拷贝
s1 := []int{1,3,6,9}
s2 := make([]int, 10)	//必须给与充足的空间
num := copy(s2, s1)

fmt.Println(s1)			//[1 3 6 9]
fmt.Println(s2)			//[1 3 6 9 0 0 0 0 0 0]
fmt.Println(num)		//4

//切片中删除元素
s1 := []int{1,3,6,9}
index := 2					//删除该位置元素
s1 = append(s1[:index], s1[index+1:]...)
fmt.Println(s1)				//[1 3 9]

// 切片拷贝
s1 := []int{1,2,3,4,5}
s2 := []int{6,7,8}
copy(s1,s2) 				//复制s2前三个元素到slice1前3位置
copy(s2,s1)	 				//复制s1前三个元素到slice2
```

注意：不会编译错误，默认第二个参数后是元素值，传入切片需要展开。如果追加的长度超过当前已分配的存储空间，切片会自动分配更大的内存。

**2.2 切片的一些简便操作**

* slice的默认开始位置是0，ar\[:n\]等价于ar\[0:n\]
* slice的第二个序列默认是数组的长度，ar\[n:\]等价于ar\[n:len\(ar\)\]
* 如果从一个数组里面直接获取slice，可以这样ar\[:\]，因为默认第一个序列是0，第二个是数组的长度，即等价于ar\[0:len\(ar\)\]
* 切片的遍历可以使用for循环，也可以使用range函数

```text
// 声明一个数组
var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
// 声明两个slice
var aSlice, bSlice []byte

// 演示一些简便操作
aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
aSlice = array[5:] // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
aSlice = array[:] // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素

// 从slice中获取slice
aSlice = array[3:7] // aSlice包含元素: d,e,f,g，len=4，cap=7
bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
bSlice = aSlice[:3] // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
bSlice = aSlice[:] // bSlice包含所有aSlice的元素: d,e,f,g
```

**2.3 切片的截取**

* `s[n]`：切片s中索引为位置为n的项
* `s[:]`：从切片s的索引位置0到`len(s)-1`所获得的切片
* `s[low:]`：从切片s的索引位置low到`len(s)-1`所获得的切片
* `s[:high]`：从切片s的索引位置0到high所获得的切片
* `s[low:high]`：从切片s的索引位置low到high所获得的切片
* `s[low:high:max]`：从low到high的切片，且容量`cap=max-low`

**1.7 字符串转切片**

```text
str := "hello,世界"
a := []byte(str)		//字符串转换为[]byte类型切片
b := []rune(str)		//字符串转换为[]rune类型切片
```

### 三 切片存储结构

与数组相比，切片多了一个存储能力值的概念，即元素个数与分配空间可以是两个不同的值，其结构如下所示：

```text
type slice struct {
	arrary = unsafe.Pointer		//指向底层数组的指针
	len int						//切片元素数量
	cap int						//底层数组的容量
}
```

所以切片通过内部的指针和相关属性引用数组片段，实现了变长方案，Slice并不是真正意义上的动态数组。

合理设置存储能力，可以大幅提升性能，比如知道最多元素个数为50，那么提前设置为50，而不是先设为30，可以明显减少重新分配内存的操作。

### 四 切片作为函数参数

```text
func test(s []int)  {
	fmt.Printf("slice address : %p",s)
	fmt.Println()
	s = append(s,4,5,6)
	fmt.Printf("after change slice address : %p",s)
	fmt.Println()
	fmt.Println(s)
}

func main()  {
	s := []int{1,2,3}
	fmt.Printf("original slice address : %p",s)
	fmt.Println()
	fmt.Println(s)
	test(s)
}

# print-------
#original slice address : 0xc0000b6000
#[1 2 3]
#slice address : 0xc0000b6000
#after change slice address : 0xc0000ac060
#[1 2 3 4 5 6]
```

```go
func test(s *[]int)  {
	fmt.Printf("slice address : %p",*s)
	fmt.Println()
	*s = append(*s,4,5,6)
	fmt.Printf("after change slice address : %p",*s)
	fmt.Println()
	fmt.Println(s)
}

func main()  {
	s := []int{1,2,3}
	fmt.Printf("original slice address : %p",s)
	fmt.Println()
	fmt.Println(s)
	test(&s)
	fmt.Printf("after change slice : %v",s)
	fmt.Println()
	fmt.Printf("after change slice address : %p",s)
}



//original slice address : 0xc00001a0a0
//[1 2 3]
//slice address : 0xc00001a0a0
//after change slice address : 0xc000018150
//&[1 2 3 4 5 6]
//after change slice : [1 2 3 4 5 6]
//after change slice address : 0xc000018150

```

