## ProtoBuf简介

Protocol Buffer的简称。Google旗下的一款平台无关，语言无关，可扩展的序列化结构数据格式，适合用于数据存储，作为不同应用、语言之间相互通信的数据交换格式,序列化后的数据为二进制数据（pb格式的数据），类比XML、JSON。

官网地址 [https://developers.google.com/protocol-buffers/](https://links.jianshu.com/go?to=https%3A%2F%2Fdevelopers.google.com%2Fprotocol-buffers%2F)

## 安装ProtoBuf编译器

从github上下载编译器源码安装包，[https://github.com/protocolbuffers/protobuf/releases](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fprotocolbuffers%2Fprotobuf%2Freleases)

## 定义一个.proto文件

student.proto文件

```
message Student
{
      required string id  = 1;
      required string name = 2;
      required string address = 3;
}
```

## 编译生成访问类文件

运行下面的命令

```
protoc --js_out=import_style=commonjs,binary:. student.proto
或者windows平台也可以如下编译
E:\portobuf>protoc.exe --js_out=import_style=commonjs,binary:. student.proto
```

会当前目录生成

```
student_pb.js
```

其中的--js_out的语法如下：

```
--js_out=[OPTIONS:]output_dir
```

如上面的例子中的option为 import_style=commonjs,binary， "."为生成文件的目录，这里为当前目录

## 打包为web可用的js文件

前置条件：需要安装npm。npm一般在安装nodejs的时候就会自动安装。

安装库文件的引用库

```
npm install -g require
```

安装打包成前端使用的js文件

```
 npm install -g browserify
```

安装protobuf的库文件

```
  npm install google-protobuf
```

打包js文件exports.js

```
  var student= require('./student_pb');
  module.exports = {
      DataProto: student
  }
```

编译生成可用js文件

```
   browserify exports.js -o  student_pb_web.js
```

## API

> 普通类型字段（required/optional）
> get{FIELD}()
> set{FIELD}(value)
> clear{FIELD}(value)
> 数组类型字段操作（repeated）
> add{FIELD}(value)
> clear{FIELD}List()
> get{FIELD}List()
> setInterestList(array)
>
> 序列化/反序列化
> serializeBinary() // 序列化
> deserializeBinary(bin) // 反序列化
>
> 调试
> toObject()

## 使用

```
    <script type="text/javascript">
        var student= new proto.Student();
        student.setId("110105199001010101");
        student.setName("沙克");
        student.setAddress("北京朝阳");
        console.log(student.toObject());
        var bytes = student.serializeBinary();
        console.log(bytes );//序列化后的数据
    </script>
```