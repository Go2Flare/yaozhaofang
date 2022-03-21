package model

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"os"
)

//初始化配置, 只要适用到viper都要先初始化
func InitConfig() {
	workDir, err := os.Getwd()
	//workDir := "D:\\My_code\\go\\src\\go_code\\yaozhaofang\\web"
	//workDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir)
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

var (
	//创建数据库连接句柄
	GlobalDB *gorm.DB
	//创建redis连接池
	GlobalRedis redis.Pool
)
//初始化redis链接
func InitRedis() {
	InitConfig()
	GlobalRedis = redis.Pool{
		MaxIdle:     20,
		MaxActive:   50,
		IdleTimeout: 60 * 5,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(viper.GetString("redis.network"),
				viper.GetString("redis.address"),
				redis.DialPassword(viper.GetString("redis.password")))},
	}
}

