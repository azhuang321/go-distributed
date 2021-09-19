# 标量数值类型

一个标量消息字段可以含有一个如下的类型——该表格展示了定义于.proto文件中的类型，以及与之对应的、在自动生成的访问类中定义的类型：

| .proto   | 说明                                                         | Python | Go      |
| -------- | ------------------------------------------------------------ | ------ | ------- |
| double   |                                                              | float  | float64 |
| float    |                                                              | float  | float32 |
| int32    | 使用变长编码，对于负值的效率很低，如果你的域有可能有负值，请使用sint64替代 | int    | int32   |
| uint32   | 使用变长编码                                                 | int    | uint32  |
| uint64   | 使用变长编码                                                 | int    | uint64  |
| sint32   | 使用变长编码，这些编码在负值时比int32高效的多                | int    | int32   |
| sint64   | 使用变长编码，有符号的整型值。编码时比通常的int64高效。      | int    | int64   |
| fixed32  | 总是4个字节，如果数值总是比总是比228大的话，这个类型会比uint32高效。 | int    | uint32  |
| fixed64  | 总是8个字节，如果数值总是比总是比256大的话，这个类型会比uint64高效。 | int    | uint64  |
| sfixed32 | 总是4个字节                                                  | int    | int32   |
| sfixed64 | 总是8个字节                                                  | int    | int64   |
| bool     |                                                              | bool   | bool    |
| string   | 一个字符串必须是UTF-8编码或者7-bit ASCII编码的文本。         | str    | string  |
| bytes    | 可能包含任意顺序的字节数据。                                 | str    | []byte  |

你可以在文章[Protocol Buffer 编码](https://developers.google.com/protocol-buffers/docs/encoding?hl=zh-cn)中，找到更多“序列化消息时各种类型如何编码”的信息。

1. 在java中，无符号32位和64位整型被表示成他们的整型对应形似，最高位被储存在标志位中。
2. 对于所有的情况，设定值会执行类型检查以确保此值是有效。

1. 64位或者无符号32位整型在解码时被表示成为ilong，但是在设置时可以使用int型值设定，在所有的情况下，值必须符合其设置其类型的要求。
2. python中string被表示成在解码时表示成str

1. Integer在64位的机器上使用，string在32位机器上使用

# 默认值



- 对于strings，默认是一个空string
- 对于bytes，默认是一个空的bytes

- 对于bools，默认是false
- 对于数值类型，默认是0

- 对于枚举，默认是第一个定义的枚举值，必须为0;
- 对于消息类型（message），域没有被设置，确切的消息是根据语言确定的，详见[generated code guide](https://developers.google.com/protocol-buffers/docs/reference/overview?hl=zh-cn)
  对于可重复域的默认值是空（通常情况下是对应语言中空列表）。
  注：对于标量消息域，一旦消息被解析，就无法判断域释放被设置为默认值（例如，例如boolean值是否被设置为false）还是根本没有被设置。你应该在定义你的消息类型时非常注意。例如，比如你不应该定义boolean的默认值false作为任何行为的触发方式。也应该注意如果一个标量消息域被设置为标志位，这个值不应该被序列化传输。
  查看[generated code guide](https://developers.google.com/protocol-buffers/docs/reference/overview?hl=zh-cn)选择你的语言的默认值的工作细节。

# option

看这个名字，就知道是选项和配置的意思，常见的选项是配置 `go_package`

```protobuf
option go_package = "./;proto";
```

现在protoc命令生成go包的时候，如果这一行没加上，会提示错误：

```protobuf
➜  proto git:(master) ✗ protoc --go_out=:. hello.proto2020/05/21 15:59:40 WARNING: Missing 'go_package' option in "hello.proto", please specify:        option go_package = ".;proto";A future release of protoc-gen-go will require this be specified.See https://developers.google.com/protocol-buffers/docs/reference/go-generated#package for more information.
```

所以，这个`go_package`和上面那个`package proto;`有啥区别呢？有点蒙啊。

我尝试这样改一下：

```protobuf
syntax = "proto3";package protoB;option go_package = ".;protoA";
```

我看下，生成的go语言包的package到底是啥？打开，生成后的go文件：

```protobuf
# vi hello.pb.gopackage protoA...
```

发现是`protoA`，说明，go的package是受`option go_package`影响的。所以，在我们没有申请这一句的时候，系统就会用proto文件的package名字，为提示，让你也加上相同的go_package名字了。

再来看一下，这个=".;proto" 是啥意思。我改一下：

```protobuf
option go_package = "./protoA";
```

执行后，发现，生成了一个`protoA`文件夹。里面的hello.pb.go文件的package也是protoA。

所以，`.;`表示的是就在本目录下的意思么？？？行吧。

再来看一下，我们改成1个绝对的路径目录：

```protobuf
option go_package = "/";
```

所以，总结一下：

```protobuf
package protoB; //这个用来设定proto文件自身的packageoption go_package = ".;protoA";  //这个用来生成的go文件package。一般情况下，可以把这2个设置成一样
```