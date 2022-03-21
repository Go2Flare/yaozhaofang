package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//创建中间件
func Test1(ctx *gin.Context){
	fmt.Println("我是第一个中间件")
	//上下文的下一步？
	ctx.Next()
	fmt.Println("model逻辑执行完了，我是第一个中间件")
}
//创建另外一种格式的中间件
func Test2() gin.HandlerFunc {
	return func(ctx *gin.Context){
		fmt.Println("我是第二个中间件")

		//ctx.Next() //就是中间件来回传递一样123->model->321
		//return  //当前中间件回来的逻辑不执行123->model->31
		ctx.Abort() //只执行到当前的中间件，下面的逻辑和中间件都不执行，12->21(没有model)
		fmt.Println("model逻辑执行完了，我是第二个中间件")
	}
}

func Test3(ctx *gin.Context) {
	fmt.Println("我是第三个中间件")
	ctx.Next()
	fmt.Println("model逻辑执行完了，我是第三个中间件")
}


func main(){
	router := gin.Default()

//	使用中间件
	router.Use(Test1)
	router.Use(Test2())
	router.Use(Test3)
	router.GET("/test1", func(ctx *gin.Context){
		fmt.Println("这里是model逻辑")
		ctx.Writer.WriteString("helloWorld！！！")
	})
	router.Run(":8778")
}