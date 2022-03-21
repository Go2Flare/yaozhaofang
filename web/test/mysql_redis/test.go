package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type TestTable struct{
	Column string `gorm:"column"`
}


var (
	GlobalDB *gorm.DB
	RedisPool redis.Pool
)

type MysqlConfig struct{
	username, password, host, port, database, charset string
}

type RedisConfig struct{
	network, address, password string
}

func InitMysql(conf MysqlConfig){
	//配置导入
	username,password,host,port,database,charset :=
		conf.username,conf.password,conf.host,
		conf.port,conf.database,conf.charset
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
		)
	db, err := gorm.Open("mysql", dsn)
	if err != nil{
		log.Printf("mysql配置连接失败，username=%v, password=%v, host=%v, port=%v, database=%v, charset=%v\n",
			username, password, host, port, database, charset)
		log.Printf("gorm.Open : %v", err)
		return
	}
	db.DB().SetMaxIdleConns(50)//连接数量
	db.DB().SetConnMaxLifetime(60*5)//conn可重用时长
	db.SingularTable(true)

	GlobalDB = db
	db.AutoMigrate(new(TestTable))
	log.Printf("mysql配置连接成功，username=%v, password=%v, host=%v, port=%v, database=%v, charset=%v\n",
		username, password, host, port, database, charset)

}

func InitRedis(conf RedisConfig){
	network, address, password :=
		conf.network, conf.address, conf.password
	RedisPool = redis.Pool{
		MaxIdle: 20,
		MaxActive: 50,
		IdleTimeout: 60 * 5,
		Dial: func()(redis.Conn, error){
			return redis.Dial(network, address, redis.DialPassword(password))
		},
	}
	conn := RedisPool.Get()
	//获取key
	ss, err := redis.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		fmt.Println("wtf???")
		log.Printf("redis配置连接失败，network=%v, address=%v, password=%v\n",
			network, address, password)
		log.Printf("redis.Dial err : %v\n", err)
		return
	}
	log.Printf("redis配置连接成功，network=%v, address=%v, password=%v\n",
		network, address, password)
	fmt.Println("redis> KEYS *")
	for _, s := range ss{
		fmt.Println(s)
	}
}

