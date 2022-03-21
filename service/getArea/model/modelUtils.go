package model

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"os"
)

//初始化配置, 从vip里取数据
func InitConfig() {
	workDir, err := os.Getwd()
	//workDir := "D:\\My_code\\go\\src\\go_code\\yaozhaofang\\web"
	//workDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/conf")
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

func InitDb() error {
	InitConfig()
	//用viper导入配置文件
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		viper.GetString("datasource.username"),
		viper.GetString("datasource.password"),
		viper.GetString("datasource.host"),
		viper.GetString("datasource.port"),
		viper.GetString("datasource.database"),
		viper.GetString("datasource.charset"),
	)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("service user connect mysql err", err)
		return err
	}
	//初始化全局句柄
	//连接池设置
	//设置初始化数据库连接数量
	db.DB().SetMaxIdleConns(50)
	db.DB().SetConnMaxLifetime(100)
	db.DB().SetConnMaxLifetime(60 * 5)

	db.SingularTable(true)

	//默认情况下表名是复数
	GlobalDB = db

	//创建表
	return db.AutoMigrate(new(User), new(House), new(Area), new(Facility), new(HouseImage), new(OrderHouse)).Error

}

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
				redis.DialPassword(viper.GetString("redis.password")))
		},
	}
}

