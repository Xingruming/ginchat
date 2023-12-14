package models

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitMySQL() {
	//自定义日志模板 打印SQL语句
	newlogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)

	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newlogger})
	fmt.Println("MySQL inited ......")
	//user := DB.Find(models.UserBasic{})
	//DB.Find(&user)
	//fmt.Println(user)

	// 迁移 schema
	DB.AutoMigrate(&Community{})
	DB.AutoMigrate(&UserBasic{})
	DB.AutoMigrate(&Message{})
	DB.AutoMigrate(&GroupBasic{})
	DB.AutoMigrate(&Contact{})
}
