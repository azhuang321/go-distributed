```sh
docker run --name mynacos -e MODE=standalone -e JVM_XMS=128m -e JVM_XMX=128m -e JVM_XMN=128m -p 8848:8848 -d nacos/nacos-server
```

阿里巴巴开源:[官方文档](https://nacos.io/zh-cn/docs/what-is-nacos.html)

