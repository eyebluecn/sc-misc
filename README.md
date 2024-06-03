
## 安装依赖
```shell
go get github.com/eyebluecn/sc-misc-idl@master
go get github.com/eyebluecn/sc-subscription-idl@master
```

如果有的时候idl无法拉去到最新版本，看看proxy配置，可能proxy有缓存。尝试使用原版后，再拉去idl.
```shell
go env -w GOPROXY=https://goproxy.io,direct
```


难解的问题：
```text
1. 包名过长，究竟应该怎么命名包名。
2. struct定义究竟放到一个文件，还是每个文件一个struct

```



一些约定：
```text
1. 错误统一约定成一种类型。
2. 有错误，一般情况都直接往外抛出。
3. 关于panic的统一处理。
4. 以来要遵循应用架构，如果一些不好处理的依赖，可以采用listener策略，例如mq中。
```


