
## 安装依赖
```shell
go get github.com/eyebluecn/sc-misc-idl@master
go get github.com/eyebluecn/sc-subscription-idl@master
```

如果有的时候idl无法拉去到最新版本，看看proxy配置，可能proxy有缓存。尝试使用原版后，再拉去idl.
```shell
go env -w GOPROXY=https://goproxy.io,direct
```



## 首次启动

### 导入数据库
首次启动需要先导入数据库。
将`docs/db/schema.sql`中的ddl语句，在你的数据库中执行即可。


### 指定启动参数
数据库连接信息在`src/common/config/mysql_config.go#getDefaultMysqlUrl()`中修改即可。

如果你希望从启动参数中配置数据库信息，可参考下面的配置：
```shell
smart.classroom.subscription "your_username:your_password@tcp(your_host:3306)/your_schema?charset=utf8mb4&parseTime=True&loc=Local"
```


## 日常启动
执行`main.go`中的main函数即可启动。



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


