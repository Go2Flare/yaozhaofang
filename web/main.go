package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go_code/yaozhaofang/web/controller"
	"go_code/yaozhaofang/web/model"
	"go_code/yaozhaofang/web/utils"
	"net/http"
)


// LoginFilter Session校验
func LoginFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//	初始化Session容器
		s := sessions.Default(ctx)
		userName := s.Get("userName")
		resp := make(map[string]interface{})
		if userName == nil {
			resp["errno"] = utils.RECODE_SESSIONERR
			resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
			fmt.Println("用户名为空")
			ctx.JSON(http.StatusOK, resp)
			ctx.Abort() //这里返回，往下不执行
		} else {
			ctx.Next() //继续向下
		}
	}
}

func main() {
	//导入配置文件
	model.InitConfig()
	// 初始化 MySQL 链接池
	model.InitDb()
	defer model.GlobalDB.Close()
	//初始化redis连接池
	model.InitRedis()
	//1.初始化路由
	router := gin.Default()

	//使用容器，现在的redis线程池为取出mysession的数据
	router.Use(sessions.Sessions("mysession", model.NewRedisStore()))

	router.Static("/home", "view")
	//router.Static("/home", "html")

	//用路由分组
	r1 := router.Group("/api/v1.0")
	{
		//路由规范
		r1.GET("areas", controller.GetArea)
		r1.GET("imagecode/:uuid", controller.GetImageCd)
		r1.GET("smscode/:mobile", controller.GetSmsCd)
		r1.POST("users", controller.PostRet)

		r1.POST("sessions", controller.PostLogin)
		r1.GET("session", controller.GetSession)

		//开始delete，才有session逻辑,下面的路由，
		//都不用校验session，直接通过中间件获取数据
		//路由过滤器，登录后才能请求下一路由的请求
		r1.Use(LoginFilter())

		r1.DELETE("session", controller.DeleteSession)
		r1.GET("user", controller.GetUserInfo)
		r1.PUT("user/name", controller.PutUserInfo)

		r1.POST("user/avatar", controller.PostAvatar)
		r1.POST("user/auth", controller.PutUserAuth)
		r1.GET("user/auth", controller.GetUserInfo)

		//获取已发布房源信息
		r1.GET("user/houses", controller.GetUserHouses)
		//发布房源
		r1.POST("houses", controller.PostHouses)
		//添加房源图片
		r1.POST("houses/:id/images", controller.PostHousesImage)
		//展示房屋详情
		r1.GET("houses/:id", controller.GetHouseInfo)
		//展示首页轮播图
		r1.GET("house/index", controller.GetIndex)
		//搜索房屋
		r1.GET("houses", controller.GetHouses)
		//下订单
		r1.POST("orders", controller.PostOrders)
		//获取订单
		r1.GET("user/orders", controller.GetUserOrder)
		//同意/拒绝订单
		r1.PUT("orders/:id/status", controller.PutOrders)
	}
	//3.运行
	router.Run(viper.GetString("server.ip") + ":" + viper.GetString("server.port"))
}
