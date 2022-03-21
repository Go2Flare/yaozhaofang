package controller

import (
	"context"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	houseMicro "go_code/yaozhaofang/web/proto/house"
	"go_code/yaozhaofang/web/utils"
	"net/http"
	"path"
)

type HouseStu struct {
	Acreage   string   `json:"acreage"`
	Address   string   `json:"address"`
	AreaId    string   `json:"area_id"`
	Beds      string   `json:"beds"`
	Capacity  string   `json:"capacity"`
	Deposit   string   `json:"deposit"`
	Facility  []string `json:"facility"`
	MaxDays   string   `json:"max_days"`
	MinDays   string   `json:"min_days"`
	Price     string   `json:"price"`
	RoomCount string   `json:"room_count"`
	Title     string   `json:"title"`
	Unit      string   `json:"unit"`
}


//获取已发布房源信息  假数据
func GetUserHouses(ctx *gin.Context) {

	/*resp := make(hash[string]interface{})

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = ""

	ctx.JSON(http.StatusOK,resp)*/

	//获取用户名
	userName := sessions.Default(ctx).Get("userName")

	/*//测试一对多查询
	//有用户名
	var userInfo model.User
	if err := model.GlobalDB.Where("name = ?",userName).Find(&userInfo).Error;err != nil {
		fmt.Println("获取当前用户信息错误",err)
	}
	//房源信息   一对多查询
	var houses []model.House

	model.GlobalDB.Model(&userInfo).Related(&houses)
	fmt.Println("11111111",houses)*/

	microClient := houseMicro.NewHouseService("go.micro.srv.house", utils.InitMicro())
	//调用远程服务
	resp, _ := microClient.GetHouseInfo(context.TODO(), &houseMicro.GetReq{UserName: userName.(string)})

	//返回数据
	ctx.JSON(http.StatusOK, resp)
}

//发布房源
func PostHouses(ctx *gin.Context) {
	//获取数据   bind数据的时候不带自动转换   c.getInt()
	var house HouseStu
	err := ctx.Bind(&house)

	//校验数据
	if err != nil {
		fmt.Println("获取数据错误", err)
		return
	}

	//获取用户名
	userName := sessions.Default(ctx).Get("userName")

	//处理数据  服务端处理
	microClient := houseMicro.NewHouseService("go.micro.srv.house", utils.InitMicro())
	//调用远程服务
	resp, _ := microClient.PubHouse(context.TODO(), &houseMicro.Request{
		Acreage:   house.Acreage,
		Address:   house.Address,
		AreaId:    house.AreaId,
		Beds:      house.Beds,
		Capacity:  house.Capacity,
		Deposit:   house.Deposit,
		Facility:  house.Facility,
		MaxDays:   house.MaxDays,
		MinDays:   house.MinDays,
		Price:     house.Price,
		RoomCount: house.RoomCount,
		Title:     house.Title,
		Unit:      house.Unit,
		UserName:  userName.(string),
	})

	//返回数据
	ctx.JSON(http.StatusOK, resp)
}

//上传房屋图片
func PostHousesImage(ctx *gin.Context) {
	//获取数据
	houseId := ctx.Param("id")
	fileHeader, err := ctx.FormFile("house_image")
	//校验数据
	if houseId == "" || err != nil {
		fmt.Println("传入数据不完整", err)
		return
	}

	//三种校验 大小,类型,防止重名  fastdfs
	if fileHeader.Size > 50000000 {
		fmt.Println("文件过大,请重新选择")
		return
	}

	fileExt := path.Ext(fileHeader.Filename)
	if fileExt != ".png" && fileExt != ".jpg" {
		fmt.Println("文件类型错误,请重新选择")
		return
	}

	//获取文件字节切片
	file, _ := fileHeader.Open()
	buf := make([]byte, fileHeader.Size)
	file.Read(buf)

	//处理数据  服务中实现
	microClient := houseMicro.NewHouseService("go.micro.srv.house", utils.InitMicro())
	//调用服务
	resp, _ := microClient.UploadHouseImg(context.TODO(), &houseMicro.ImgReq{
		HouseId: houseId,
		ImgData: buf,
		FileExt: fileExt,
	})

	//返回数据
	ctx.JSON(http.StatusOK, resp)
}

//获取房屋详情
func GetHouseInfo(ctx *gin.Context) {
	//获取数据
	houseId := ctx.Param("id")
	//校验数据
	if houseId == "" {
		fmt.Println("获取数据错误")
		return
	}
	userName := sessions.Default(ctx).Get("userName")
	//处理数据
	microClient := houseMicro.NewHouseService("go.micro.srv.house", utils.InitMicro())
	//调用远程服务
	resp, _ := microClient.GetHouseDetail(context.TODO(), &houseMicro.DetailReq{
		HouseId:  houseId,
		UserName: userName.(string),
	})

	//返回数据
	ctx.JSON(http.StatusOK, resp)
}

func GetIndex(ctx *gin.Context) {
	//处理数据
	microClient := houseMicro.NewHouseService("go.micro.srv.house", utils.InitMicro())
	//调用服务
	resp, _ := microClient.GetIndexHouse(context.TODO(), &houseMicro.IndexReq{})

	ctx.JSON(http.StatusOK, resp)
}

//搜索房屋
func GetHouses(ctx *gin.Context) {
	//获取数据
	//areaId
	aid := ctx.Query("aid")
	//start day
	sd := ctx.Query("sd")
	//end day
	ed := ctx.Query("ed")
	//排序方式
	sk := ctx.Query("sk")
	//page  第几页
	//ctx.Query("p")
	//校验数据
	if aid == "" || sd == "" || ed == "" || sk == "" {
		fmt.Println("传入数据不完整")
		return
	}

	//处理数据   服务端  把字符串转换为时间格式,使用函数time.Parse()  第一个参数是转换模板,需要转换的二字符串,两者格式一致
	//edTime ,_:=time.Parse("2006-01-02 15:04:05",sd+" 00:00:00")
	//edTime,_ := time.Parse("2006-01-02",ed)
	//
	//sdTime,_ :=time.Parse("2006-01-02",sd)
	//edTime,_ := time.Parse("2006-01-02",ed)
	//d := edTime.Sub(sdTime)
	//fmt.Println(d.Hours())

	microClient := houseMicro.NewHouseService("go.micro.srv.house", utils.InitMicro())
	////调用远程服务
	resp, _ := microClient.SearchHouse(context.TODO(), &houseMicro.SearchReq{
		Aid: aid,
		Sd:  sd,
		Ed:  ed,
		Sk:  sk,
	})

	//返回数据
	ctx.JSON(http.StatusOK, resp)

}

