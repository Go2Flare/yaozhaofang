package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_code/yaozhaofang/web/proto/getArea"
	"go_code/yaozhaofang/web/proto/getCaptcha"
	"go_code/yaozhaofang/web/proto/register"
	"go_code/yaozhaofang/web/proto/user"
	"go_code/yaozhaofang/web/utils"
	"image/png"
	"log"
	"net/http"
	"regexp"
)

// 获取 Session 数据
func GetSession(ctx *gin.Context) {
	resp := make(map[string]interface{})

	s := sessions.Default(ctx) // 初始化 Session 对象
	//从session数据里找userName的数据
	userName := s.Get("userName")

	// 用户没有登录.---没存在 MySQL中, 也没存在 Session 中
	if userName == nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	} else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

		var nameData struct {
			Name string `json:"name"`
		}
		nameData.Name = userName.(string) // 类型断言
		fmt.Println(nameData)
		resp["data"] = nameData
	}

	ctx.JSON(http.StatusOK, resp)
}

//获取验证码图片信息，从consul里调用微服务
func GetImageCd(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	consulService := utils.InitMicro()
	//调用对应service的Call函数
	//我们这边看到的是NewGetCaptchaService(服务名，服务发现)
	fmt.Println("consulService.String : ",consulService.String())
	fmt.Println("consulService.Options : ", consulService.Options())
	fmt.Println("consulService.Addr : ", consulService.Options())

	microClient := getCaptcha.NewGetCaptchaService("go.micro.srv.getCaptcha", consulService)

	//客户端初始化了，可以调用getCaptcha函数的方法
	resp, err := microClient.MicroGetCaptcha(context.TODO(), &getCaptcha.Request{Uuid: uuid})
	if err != nil {
		fmt.Println("microClient.Cal err:", err)
	}

	//用captcha模块的图片
	var img captcha.Image
	//拿到字节流，反序列化成图片
	err = json.Unmarshal(resp.Img, &img)
	if err != nil{
		log.Printf("json.Unmarshal err : %v", err)
	}
	//将图片解码到浏览器
	png.Encode(ctx.Writer, img)
}

func GetSmsCd(ctx *gin.Context) {
	//获取手机号
	mobile := ctx.Param("mobile")
	//拆分GET请求中的URL格式==格式：资源路径?k=v&k=v&k=v
	text := ctx.Query("text")
	uuid := ctx.Query("id")

	//校验手机号格式
	reg, _ := regexp.Compile(`^1[3,4,5,7,8]\d{9}$`)
	isRightMobile := reg.MatchString(mobile)
	if !isRightMobile {
		log.Fatal("手机号格式错误")
		return
	}

	if mobile == "" || text == "" || uuid == "" {
		log.Fatal("GetSmsCd传入数据不完整")
		return
	}

	consulService := utils.InitMicro()

	//初始化客户端
	microClient := register.NewRegisterService("go.micro.srv.register", consulService)

	//调用远程函数
	resp, err := microClient.SmsCode(context.TODO(), &register.Request{
		Uuid:   uuid,
		Text:   text,
		Mobile: mobile,
	})

	if err != nil {
		fmt.Println("microClient.SendSms err:", err)
		return
	}


	//写入校验结果在浏览器
	ctx.JSON(http.StatusOK, resp)
}

//发送注册信息
func PostRet(ctx *gin.Context) {
	type RegisterUser struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
		SmsCode  string `json:"sms_code"`
	}
	//确定容器
	resp := make(map[string]interface{})
	//绑定数据
	var regUser RegisterUser
	err := ctx.Bind(&regUser)
	if err != nil {
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(200, resp)
		return
	}
	//校验数据
	if regUser.Mobile == "" || regUser.Password == "" || regUser.SmsCode == "" {
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(200, resp)
		return
	}
	//正则校验手机号
	reg, _ := regexp.Compile(`^1[3,4,5,7,8]\d{9}$`)
	if !reg.MatchString(regUser.Mobile) {
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(200, resp)
		return
	}
	//远程调用
	//初始化客户端
	consulService := utils.InitMicro()

	//初始化客户端
	microClient := register.NewRegisterService("go.micro.srv.register", consulService)

	//调用远程函数
	response, err := microClient.Register(context.TODO(), &register.RegRequest{
		Mobile:   regUser.Mobile,
		SmsCode:  regUser.SmsCode,
		Password: regUser.Password,
	})
	fmt.Printf("当前注册的响应 response=%v, err=%v, utils.RECODE_OK=%v\n",response, err,utils.RECODE_OK)
	if err != nil {
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if response.Errno == utils.RECODE_OK {
		s := sessions.Default(ctx)
		s.Set("userName", regUser.Mobile)
		s.Save()
		fmt.Println("注册流程结束")
	}
}

// 测试实现：
// 获取地域信息的微服务
func GetArea(ctx *gin.Context) {

	//调用远程逻辑获取地域信息
	microClient := getArea.NewGetAreaService("go.micro.srv.getArea", utils.InitMicro())

	//连接传参
	resp, err := microClient.MicroGetArea(context.TODO(), &getArea.Request{})
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, resp)

}

//登录逻辑调用微服务
func PostLogin(ctx *gin.Context) {
	//	获取前端数据
	var loginData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
	}
	err := ctx.Bind(&loginData)
	if err != nil {
		fmt.Println("PostLogin获取数据失败")
		return
	}

	//初始化客户端
	//把登录放在微服务了
	microClient := register.NewRegisterService("go.micro.srv.register", utils.InitMicro())

	//调用远程服务
	resp, err := microClient.Login(context.TODO(), &register.RegRequest{
		Mobile:   loginData.Mobile,
		Password: loginData.PassWord,
	})
	defer ctx.JSON(http.StatusOK, resp)
	if err != nil {
		fmt.Println("调用login服务错误", err)
		return
	}
	//返回数据  存储session  并返回数据给web端
	session := sessions.Default(ctx)
	session.Set("userName", resp.Name)
	session.Save()

}

// 退出登录不用微服务,正常逻辑删除session即可
func DeleteSession(ctx *gin.Context) {
	resp := make(map[string]interface{})

	// 初始化 Session 对象
	s := sessions.Default(ctx)
	// 删除 Session 数据
	s.Delete("userName") // 没有返回值
	// 必须使用 Save 保存
	err := s.Save() // 有返回值

	if err != nil {
		resp["errno"] = utils.RECODE_IOERR // 没有合适错误,使用 IO 错误!
		resp["errmsg"] = utils.RecodeText(utils.RECODE_IOERR)

	} else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	}
	ctx.JSON(http.StatusOK, resp)
}

// 获取用户基本信息,调用user/MicroGetUser微服务实现
func GetUserInfo(ctx *gin.Context) {
	s := sessions.Default(ctx)    // Session 初始化
	userName := s.Get("userName") // 根据key 获取Session

	//调用user/MicroGetUser微服务
	microClient := user.NewUserService("go.micro.srv.user", utils.InitMicro())

	resp, err := microClient.MicroGetUser(context.TODO(), &user.Request{
		Name: userName.(string),
	})

	if err != nil {
		fmt.Println("web/controller/user GetUserInfo调用微服务出错", err)
		resp.Errno = utils.RECODE_DATAERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
	}

	ctx.JSON(http.StatusOK, resp)

}

//修改用户名
func PutUserInfo(ctx *gin.Context) {
	//	获取当前的用户名
	s := sessions.Default(ctx) //初始化session对象
	userName := s.Get("userName")

	//	获取新的用户名
	var nameData struct {
		Name string `json:"name"`
	}
	err := ctx.Bind(&nameData)
	if err != nil {
		fmt.Println("web/controller/user PutUserInfo 获取数据 err", err)
	}

	//调用微服务UpdateUserName
	microClient := user.NewUserService("go.micro.srv.user", utils.InitMicro())

	//调用远程服务
	resp, _ := microClient.UpdateUserName(context.TODO(), &user.UpdateReq{
		NewName: nameData.Name,
		OldName: userName.(string),
	})

	//更新session数据
	if resp.Errno == utils.RECODE_OK {
		s.Set("userName", nameData.Name)
		s.Save()
	}
	ctx.JSON(http.StatusOK, resp)
}

//校验用户真实姓名
func PutUserAuth(ctx *gin.Context) {
	//	获取用户数据
	type AuthStu struct {
		IdCard   string `json:"id_card"`
		RealName string `json:"real_name"`
	}
	var auth AuthStu
	err := ctx.Bind(&auth)
	//	校验数据
	if err != nil {
		fmt.Println("获取数据错误", err)
		return
	}

	//取出session数据
	session := sessions.Default(ctx)

	userName := session.Get("userName")

	//	处理数据的微服务
	microClient := user.NewUserService("go.micro.srv.user", utils.InitMicro())
	//调用远程服务
	resp, _ := microClient.AuthUpdate(context.TODO(), &user.AuthReq{
		UserName: userName.(string),
		RealName: auth.RealName,
		IdCard:   auth.IdCard,
	})
	ctx.JSON(http.StatusOK, resp)
}
