package model

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	ginRedis "github.com/gin-contrib/sessions/redis"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"log"
	"testing"
)

func TestInitConfig(t *testing.T){
	InitConfig()
}

func TestInitRedis(t *testing.T){
	InitRedis()
}

func TestRedis(t *testing.T){
	InitRedis()
	//初始化一条连接
	conn := GlobalRedis.Get()
	bytes, err := redis.Bytes(conn.Do("SET", "yoyo", "shit"))
	fmt.Println("存储得到的redis响应 ：", string(bytes))
	bytes, err = redis.Bytes(conn.Do("GET", "yoyo"))
	fmt.Println("get的数据为 ：", string(bytes))
	if err != nil{
		log.Fatalf("err : %v", err)
	}
	store, _ := ginRedis.NewStore(10, viper.GetString("redis.network"),
		viper.GetString("redis.address"),
		viper.GetString("redis.password"), []byte("itcast"))

	//	使用容器
	sessions.Sessions("mysession", store)
}