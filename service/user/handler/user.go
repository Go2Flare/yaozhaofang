package handler

import (
	"context"
	"fmt"
	"github.com/tedcy/fdfs_client"
	"go_code/yaozhaofang/service/user/model"
	"go_code/yaozhaofang/service/user/proto/user"
	"go_code/yaozhaofang/service/user/utils"
	"os"
)

type User struct{}


//根据用户名获取用户信息 在mysql数据库中查找
// Call is a single request handler called via client.Call or the generated client code
func (e *User) MicroGetUser(ctx context.Context, req *user.Request, rsp *user.Response) error {
	if req.Name == "" { // 用户没登录, 但进入该页面, 恶意进入.
		rsp.Errno = utils.RECODE_SESSIONERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_SESSIONERR)
	}

	myUser, err := model.GetUserInfo(req.Name)
	if err != nil {
		rsp.Errno = utils.RECODE_USERERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_USERERR)
		return err
	}
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	//获取一个结构体对象
	var userInfo user.UserInfo
	userInfo.UserId = int32(myUser.ID)
	userInfo.Name = myUser.Name
	userInfo.Mobile = myUser.Mobile
	userInfo.RealName = myUser.Real_name
	userInfo.IdCard = myUser.Id_card
	userInfo.AvatarUrl = "http://47.106.87.191:8800/" + myUser.Avatar_url

	rsp.Data = &userInfo

	return nil
}

//修改用户名
func (e *User) UpdateUserName(ctx context.Context, req *user.UpdateReq, resp *user.UpdateResp) error {
	//根据传递过来的用户名更新数据中新的用户名
	err := model.UpdateUserName(req.OldName, req.NewName)
	if err != nil {
		fmt.Println("更新失败", err)
		resp.Errno = utils.RECODE_DATAERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		//micro规定如果有错误,服务端只给客户端返回错误信息,不返回resp,如果没有错误,就返回resp
		return nil
	}

	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	var nameData user.NameData
	nameData.Name = req.NewName

	resp.Data = &nameData

	return nil
}

//1.上传用户头像1.0
func (e *User) UploadAvatar(ctx context.Context, req *user.UploadReq, resp *user.UploadResp) error {

	workDir, _ := os.Getwd()
	//存入到fastdfs中
	fClient, err := fdfs_client.NewClientWithConfig(workDir + "/conf/fdfs.conf")
	if err != nil {
		fmt.Println("service UploadAvatar user save image to fastdfs err", err)
		return err
	}

	//上传文件到fdfs
	fdfsResp, err := fClient.UploadByBuffer(req.Avatar, req.FileExt[1:])
	if err != nil {
		fmt.Println("图片上传错误", err)
		resp.Errno = utils.RECODE_DATAERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return err
	}

	//把存储凭证写入数据库
	err = model.UpdateAvatar(req.UserName, fdfsResp)
	if err != nil {
		fmt.Println("存储用户头像错误", err)
		resp.Errno = utils.RECODE_DBERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return nil
	}

	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	var uploadData user.UploadData
	uploadData.AvatarUrl = "http://47.106.87.191:8866/" + fdfsResp
	resp.Data = &uploadData
	return nil
}

//调用借口校验realName和idcard是否匹配
func (e *User) AuthUpdate(ctx context.Context, req *user.AuthReq, resp *user.AuthResp) error {
	//调用借口校验realName和idcard是否匹配

	//存储真实姓名和真是身份证号  数据库
	err := model.SaveRealName(req.UserName, req.RealName, req.IdCard)
	if err != nil {
		resp.Errno = utils.RECODE_DBERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return nil
	}

	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	return nil
}
