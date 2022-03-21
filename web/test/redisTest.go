package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	redisUtils "go_code/yaozhaofang/web/test/redis"
)

func main() {
	redisUtils.InitConfig()
	fmt.Println("viper", viper.GetString("redis.network"),
		viper.GetString("redis.address"),
		viper.GetString("redis.password"))
	//	1.连接数据库
	//	conn, err := redis.Dial(viper.GetString("redis.network"), viper.GetString("redis.address"), redis.DialPassword(viper.GetString("redis.password")))
		conn, err := redis.Dial("tcp", "localhost:6379", redis.DialPassword("4.234.23123"))
	if err!=nil{
		fmt.Println("redis Dial err", err)
	}
	//defer conn.Close()
	//	2.操作数据库
	//	推荐用Do()函数
	reply, err := conn.Do("set", "flare", "owner")
	r,e := redis.String(reply, err)
	fmt.Println(r, e)
	conn.Close()
}
