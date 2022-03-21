package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router:= gin.Default()

	router.GET("/test", func(context *gin.Context){
	////	设置cookie
	//	context.SetCookie("mytest", "nmlgb", 0, "", "", false, true)
	//	context.Writer.WriteString("真的老牛逼了。。。cookie测试成功")

//	获取cookie
	cookieVal, _ := context.Cookie("mytest")
	fmt.Println("获取的cookie为",cookieVal)
	context.Writer.WriteString("测试Cookie")
})

	router.Run(":52669")
}
