# 数据结构概述

### 一 数据结构概念

在实际开发中，往往需要将许多数据分门别类处理，我们需要为这些数据开辟存储空间进行统一存储。不同类数据往往存储的要求不一致，当数据存储进一个变量后，有些需求要求能够快速查找出某个单独的数据，有些需求要求能够快速存储，根据这些需求而生成的不同数据存储容器我们称之为数据结构。

> 数据结构 data structure：相互之间存在一种或多种特定关系的数据集。 开发中常见的数据结构有：数组、双向链表、二叉树等等。

数据结构是一个二元组（多元组用来描述确定成分的数学对象，即对象个数优先的序列）：

```text
# 表示方式   D是数据元素的有限集合，S是在D中数据元素之间的关系集合；<x,y> 表示数据元素与y之间有关系
Data_Structure = (D, S)                     

# 线性结构示例
D = {01, 02, 03, 04, 05}，S = {<02,04>, <03,05>, <05,02>, <01,03>}

# 树结构示例
D = {01, 02, 03, 04, 05, 06}，S = {<01,02>, <01,03>, <02,04>, <02,05>, <03,06>}
```

### 二 结构与存储

#### 2.1 逻辑结构

> 逻辑结构：数据结构实例中的数据元素之间存在的相互关系

常见的逻辑结构有四种：

* 集合结构：结构中的数据元素之间除了同属于一个集合外，无其他关系，如并查集
* 线性结构：结构中的数据元素之间存在着一对一的关系，如线性表、向量、栈、队列、优先队列、字典
* 树形结构：结构中的数据元素之间存在着一对多的关系，如二叉树
* 图形结构：结构中的数据元素之间存在着多对多的关系，也称为网状结构，如有向图

**集合结构**：元素之间完全平等，只有一个关系，即属于同一集合，如下图：

![](https://github.com/overnote/over-algorithm/raw/master/images/structure/01-01.svg)

**线性结构**：元素之间是一对一关系，包括数组、链表等常见数据结构，如下图所示：

![](https://github.com/overnote/over-algorithm/raw/master/images/structure/01-02.svg)

**树形结构**：元素之间是一对多关系，常见的数据结构有二叉树，如下图所示：

![](https://github.com/overnote/over-algorithm/raw/master/images/structure/01-03.svg)

**图形结构**：元素之间是多对多关系，常见的数据结构是图，如下图所示：

![](https://github.com/overnote/over-algorithm/raw/master/images/structure/01-04.svg)

整体而言，数据结构的逻辑结构可以划分为

* 线性结构
* 非线性结构：树、图

线性结构具备下列特点：

* 必定存在唯一的一个 “第一个元素”
* 必定存在唯一的一个 “最后一个元素”
* 除第一个元素之外，其他数据元素均有唯一的前驱

#### 2.2 物理结构

> 物理结构（存储结构）：数据的逻辑结构是在计算机中的真实存储形式，即数据结构在计算机中的表示（映像），也称为存储结构。

物理结构既包括元素本身的表示，也包括元素关系的表示。2.1中的逻辑结构是在物理结构中实现的，也就是说物理结构既表示了数据元素，也表示了数据元素的关系！

如果逻辑结构不能在物理中实现，数据结构也就失去了意义！

一些常识：

```text
位/bit：是计算机中的最小单元，即二进制数的一位。

数据元素可以由若干位组合形成的位串表示，这个位串称为**元素**（element）或**结点**（node）。

当数据元素由若干数据项组成时，位串中对应于各个数据项的子位串称为**数据域**（data field）。  

所以结点也可以看作是数据元素在计算机中的映像。  
```

物理结构（元素的关系）在计算机中有两种不同的表示方法：

* 顺序映像，其存储结构称为顺序存储结构，特点是借助元素在存储器中的相对位置来表示数据元素之间的逻辑关系
* 非顺序映像，其存储结构称为链式存储结构，特点是利用元素存储地址的指针表示 数据元素之间的逻辑关系

**顺序结构**：把数据元素存放在连续的存储单元里，其数据间的逻辑关系和物理关系是一致的，最经典的顺序结构是数组，数组中的元素都是依次摆放的，如下所示：

[![](https://github.com/overnote/over-algorithm/raw/master/images/structure/01-05.svg)](https://github.com/overnote/over-algorithm/blob/master/images/structure/01-05.svg)

顺序结构无法解决插队等问题，很多数据要求存储的结构具有变化性。比如在银行办理业务时，按顺序结构领了排队号码，但是在等待期间，你自己是可以随处走动的。那么形容人在等待时期的一系列动作数据的存储，就需要链式结构。

**链式结构**：

[![](https://github.com/overnote/over-algorithm/raw/master/images/structure/01-06.svg)](https://github.com/overnote/over-algorithm/blob/master/images/structure/01-06.svg)

上述两种存储结构，在具体存储时，可以采取4种存储方法：

* 顺序存储方法：即顺序结构的存储方式
* 链式存储方法：即链式结构的存储方式
* 索引存储方法：除了要存储节点信息外，还建立附加的索引来标识节点的地址，索引项常常是 k,v 出现
* 散列存储方法：根据结点的关键字，通过散列函数计算出该节点的存储地址，本质上是顺序存储的扩展

#### 2.3 物理结构和逻辑结构关系

逻辑结构是面向问题的：实际业务中，我们需要什么样的数据结构，要根据业务出发选择合适的结构。

物理结构是面向计算机的：选型了逻辑结构后，还需要将这些数据按照逻辑结构规范存储进计算机中。

任何一个算法的设计取决于逻辑结构，其实现依赖于存储结构。

#### 三 抽象数据类型

抽象数据类型ADT（Abstract Data Type\)：

> 数据结构的数学数据模型以及定义在该模型上的一系列操作。  
> 抽象数据类型由三元组表示 \(D, S, P\)，D是数据对象，S是D上的关系集，P是对D的操作集

ADT定义格式如下：

```text
ADT 抽象数据类型名 {
    数据对象：<数据对象的定义>
    数据关系：<数据关系的定义>
    基本操作：<基本操作的定义>
}ADT 抽象数据类型名
```

严蔚敏《数据结构》中对线性表的抽象数据类型定义：

```text
ADT List {
    数据对象：D = {a1 | a1 ∈ ElemSet, i = 1, 2, ...,n, n >= 0}
    数据关系：R1 = { <ai-1, ai > | ai-1, ai∈D， i=2,...n}           # 类似ai-1中i-1都是下标
    基本操作：
        InitList(&L)
            操作结果：构造一个空线性表L
        DestroyLit(&L)
            初始条件：线性表L已经存在
            操作结果：销毁线性表L
        ClearList(&L)
            初始条件：线性表L已经存在
            操作结果：重置L为空表
        ListEmpty(L)
            初始条件：线性表L已经存在
            操作结果：若L为空表，返回TRUE，否则返回FALSE
        ListLength(L)
            初始条件：线性表L已经存在
            操作结果：返回L中数据元素的个数
        GetElem(L, e, compare())
            初始条件：线性表L已经存在，1<=i<=ListLenth(L)
            操作结果：用e返回L中第i个数据元素的值
        LocateElem(L, e, compare())
            初始条件：线性表L已经存在，compare()是 数据元素判定函数
            操作结果：返回L中第1个与e满足关系compare()的数据元素位序，若这样的数据元素不存在，则返回0
        ProrElem(L, cur_e, &pre_e)
            初始条件：线性表L已存在
            操作结果：若cur_e是L的数据元素，且不是最后一个，则用pre_e返回它的前驱，否则操作失败，pre_e无定义
        NextElem(L, cur_e, &next)e)
            初始条件：线性表L已经存在
            操作结果：若cure_e是L的数据元素，且不是最后一个，则用next_e返回它的后继，否则操作失败，next_e无定义
        ListInsert(&L, i, e)
            初始条件：线性表L已经存在，1<=i<=ListLength(L)+1
            操作结果：在L中第i个位置之前插入新的数据元素e，L的长度加1
        ListDelete(&L, i, &e)
            初始条件：线性表L已经存在且非空，1<=i<=ListLength(L)
            操作结果：删除L的第i个数据元素，并用e返回其值，L的长度减1
        ListTraverse(L, visit())
            初始条件：线性表L已经存在
            操作结果：依次对L的每个数据元素调用函数visit()，一旦visit()失败，则操作失败
}ADT List
```

