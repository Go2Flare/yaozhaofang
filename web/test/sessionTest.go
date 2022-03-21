package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"log"
)
func main() {
	router := gin.Default()

//	gin框架的连接redis句柄，可以将session存储在redis后台，后面加密的串可以随便写
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "4.234.23123", []byte("abcdefg"))
	if err != nil {
		log.Fatalf("redis.NewStore err := %v\n",err)
	}

	//使用session，调用session中间件（session命名，redis的连接句柄）
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/test", func(context *gin.Context){
	//	调用session, 设置session数据
		s:=sessions.Default(context)
		//第一次获取，设置session
		//s.Set(11111, 22222)
		//s.Save()
		//--------
		//第二次获取，直接get session
		v:= s.Get(11111)
		fmt.Println("获取的Session",v.(int))
		context.Writer.WriteString("测试Session....")
	})

	router.Run(":5266")
}