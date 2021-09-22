# 1.安装

```shell
docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600/udp --name myconsul consul consul agent -dev -client=0.0.0.0
#开启docker服务时,启动 此容器
docker container update --restart=always myconsul
```

默认端口打开
```sh
http://wslhost:8500/ui/dc1/services
```

测试dns,后期网关服务要用

```sh
dig @127.0.0.1 -p 8600 consul.service.consul SRV
```

[官网](https://www.consul.io/)

