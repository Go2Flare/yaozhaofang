package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_code/yaozhaofang/web/model"
	"go_code/yaozhaofang/web/utils"
	"net/http"
)

func getUserInfo(ctx *gin.Context){
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

//	session初始化
	s := sessions.Default(ctx)
	userName := s.Get("userName")
	if userName == nil{
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SERVERERR)
		return
	}
	//获取用户信息
	user ,err := model.GetUserInfo(userName.(string))
	if err !=nil{
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		return
	}
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	//接收用户信息
	temp := make(map[string]interface{})
	temp["user_id"] = user.ID
	temp["name"] = user.Name
	temp["mobile"] = user.Mobile
	temp["real_name"] = user.Real_name
	temp["id_card"] = user.Id_card
	temp["avatar_url"] = user.Avatar_url
	resp["data"] = temp

}