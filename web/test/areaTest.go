package main

import (
	"fmt"
	"go_code/yaozhaofang/web/model"
	"go_code/yaozhaofang/web/utils"
)

//获取mysql句柄
func main() {
	model.InitRedis()
	model.InitDb()
	defer model.GlobalDB.Close()
	//	mysql中获取数据
	GetArea()
}

func GetArea() {

	var areas []model.Area

	model.GlobalDB.Find(&areas)

	//	数据写入redis中

	conn := model.GlobalRedis.Get()
	conn.Do("set", "areaData", areas)

	resp := make(map[string]interface{})

	resp["errno"] = "0"
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = areas

	fmt.Println("数据写入redis。。")
}

//func GetArea() {
//	var areas []model.Area
//	model.GlobalConn.Find(&areas)
//	conn:= model.RedisPool.Get()
//	conn.Do("set","areaData",areas)
//	resp:=make(hash[string]interface{})
//	resp["errno"] = "0"
//	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
//	resp["data"] = areas
//}
