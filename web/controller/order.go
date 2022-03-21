package controller

import (
	"context"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	orderMicro "go_code/yaozhaofang/web/proto/userOrder"
	"go_code/yaozhaofang/web/utils"
	"net/http"
)

//用户订单
type OrderStu struct {
	EndDate   string `json:"end_date"`
	HouseId   string `json:"house_id"`
	StartDate string `json:"start_date"`
}

//下订单
func PostOrders(ctx *gin.Context) {
	//获取数据
	var order OrderStu
	err := ctx.Bind(&order)

	//校验数据
	if err != nil {
		fmt.Println("获取数据错误", err)
		return
	}
	//获取用户名
	userName := sessions.Default(ctx).Get("userName")

	//处理数据  服务端处理业务
	microClient := orderMicro.NewUserOrderService("go.micro.srv.userOrder", utils.InitMicro())
	//调用服务
	resp, _ := microClient.CreateOrder(context.TODO(), &orderMicro.Request{
		StartDate: order.StartDate,
		EndDate:   order.EndDate,
		HouseId:   order.HouseId,
		UserName:  userName.(string),
	})

	//返回数据
	ctx.JSON(http.StatusOK, resp)
}

//获取订单信息
func GetUserOrder(ctx *gin.Context) {
	//获取get请求传参
	role := ctx.Query("role")
	//校验数据
	if role == "" {
		fmt.Println("获取数据失败")
		return
	}

	//处理数据  服务端
	microClient := orderMicro.NewUserOrderService("go.micro.srv.userOrder", utils.InitMicro())
	//调用远程服务
	resp, _ := microClient.GetOrderInfo(context.TODO(), &orderMicro.GetReq{
		Role:     role,
		UserName: sessions.Default(ctx).Get("userName").(string),
	})

	//返回数据
	ctx.JSON(http.StatusOK, resp)
}

type StatusStu struct {
	Action string `json:"action"`
	Reason string `json:"reason"`
}

//更新订单状态
func PutOrders(ctx *gin.Context) {
	//获取数据
	id := ctx.Param("id")
	var statusStu StatusStu
	err := ctx.Bind(&statusStu)

	//校验数据
	if err != nil || id == "" {
		fmt.Println("获取数据错误", err)
		return
	}

	//处理数据   更新订单状态
	microClient := orderMicro.NewUserOrderService("go.micro.srv.userOrder", utils.InitMicro())
	//调用元和产能服务
	resp, _ := microClient.UpdateStatus(context.TODO(), &orderMicro.UpdateReq{
		Action: statusStu.Action,
		Reason: statusStu.Reason,
		Id:     id,
	})

	//返回数据
	ctx.JSON(http.StatusOK, resp)
}