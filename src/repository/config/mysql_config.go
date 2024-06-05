package config

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

var (
	DB *gorm.DB
)

type MysqlConfig struct {
}

var (
	mysqlConfig     *MysqlConfig
	mysqlConfigOnce sync.Once
)

// 采用单例模型
func DefaultMysqlConfig() *MysqlConfig {
	mysqlConfigOnce.Do(func() {
		mysqlConfig = &MysqlConfig{}
	})
	return mysqlConfig
}

// 如果启动参数没有指定mysqlUrl，那么就使用这里的参数配置。
func (receiver *MysqlConfig) getDefaultMysqlUrl() string {

	var mysqlPort int = 3306
	var mysqlHost string = "127.0.0.1"
	var mysqlSchema string = "smart_classroom_go"
	var mysqlUsername string = "smart_classroom_go"
	var mysqlPassword string = "your_db_password" //数据库密码
	var mysqlCharset string = "utf8mb4"

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlSchema, mysqlCharset)
}

// 从启动参数中解析出mysqlUrl.
func (receiver *MysqlConfig) getMysqlUrl() string {

	var mysqlUrl string = receiver.getDefaultMysqlUrl()

	if len(os.Args) >= 2 {
		mysqlUrl = os.Args[1]
		klog.Infof("从启动参数读取mysql url: %s", mysqlUrl)
	}

	return mysqlUrl
}

func (receiver *MysqlConfig) Init() *gorm.DB {

	ctx := context.Background()
	mysqlUrl := receiver.getMysqlUrl()
	klog.CtxInfof(ctx, "MySQL URL: %v", mysqlUrl)

	//log strategy.
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // slow SQL 1s
			LogLevel:                  logger.Warn, // log level. open when debug.
			IgnoreRecordNotFoundError: true,        // ignore ErrRecordNotFound
			Colorful:                  false,       // colorful print
		},
	)

	db, err := gorm.Open(mysql.Open(mysqlUrl), &gorm.Config{Logger: dbLogger})
	if err != nil {
		klog.CtxInfof(ctx, "MySQL 连接失败: %v", err.Error())
		panic(err)
	}
	DB = db
	return DB
}
