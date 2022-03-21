package controller

import (
	"context"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_code/yaozhaofang/web/proto/user"
	"go_code/yaozhaofang/web/utils"
	"net/http"
	"path"
)

//1.上传图片调用微服务
func PostAvatar(ctx *gin.Context) {
	//以下调用user微服务的uploadAvatar-------------------------------------

	// 获取图片文件, 静态文件对象
	file, _ := ctx.FormFile("avatar")
	// 上传文件到项目中
	//err := ctx.SaveUploadedFile(file, "test/"+file.Filename)
	//fmt.Println("上传图片结果", err)

	//上传文件到fastDFS中
	//clt, err := fdfs_client.NewClientWithConfig("D:\\My_code\\go\\src\\go_code\\yaozhaofang\\web\\conf\\fdfs.conf")
	//
	//打开文件读取字节流
	f, err := file.Open() //只读打开
	if err != nil {
		fmt.Println("文件打开错误，err", err)
		return
	}

	//按文件实际大小打开文件,调用要传的
	buf := make([]byte, file.Size)
	//读取文件内容
	f.Read(buf)

	//	go根据文件名获取文件后缀,调用要传的长度参数
	fileExt := path.Ext(file.Filename) //默认带有.号

	//	按字节流上传，去除. rpc调用传入字节流和长度
	//remoteID, err := clt.UploadByBuffer(buf, fileExt[1:])
	//if err != nil {
	//	fmt.Println("图片上传错误", err)
	//	return
	//}
	//fmt.Println(remoteID)

	//获取session，得到当前用户
	userName := sessions.Default(ctx).Get("userName")

	//微服务调用
	microClient := user.NewUserService("go.micro.srv.user", utils.InitMicro())

	//调用远程函数
	resp, _ := microClient.UploadAvatar(context.TODO(), &user.UploadReq{
		UserName: userName.(string),
		Avatar:   buf,
		FileExt:  fileExt,
	})

	//根据用户名，更新用户头像，--mysql数据库，存入数据库只要id就行，
	//不要加http头，如果以后换服务器，就得更新库就麻烦了
	//	return GlobalConn.Model(new(User)).Where("name = ?", userName).Update("avatar_url", avatar).Error
	//model.UpdateAvatar(userName.(string), remoteID)

	//响应请求
	//resp := make(hash[string]interface{})
	//resp["errno"] = utils.RECODE_OK
	//resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	//
	//temp := make(hash[string]interface{})
	//temp["avatar_url"] = "http://47.106.87.191:8800/" + remoteID
	//resp["data"] = temp

	ctx.JSON(http.StatusOK, resp)
}
