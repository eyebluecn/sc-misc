
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
数据库连接信息在`src/repository/config/mysql_config.go` 中修改即可。

如果你希望从启动参数中配置数据库信息，可参考下面的配置：
```shell
smart.classroom.subscription "your_username:your_password@tcp(your_host:3306)/your_schema?charset=utf8mb4&parseTime=True&loc=Local"
```

### 代码生成
本项目中使用了gorm，可以自动生成数据库模型和基础查询代码，在命令行执行 `script/gen-gorm.sh` 即可，产出物放置于
`model/po`和`repository/query`。


## 日常启动
执行`main.go`中的main函数即可启动。


