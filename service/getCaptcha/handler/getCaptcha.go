package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/micro/go-micro/util/log"
	"go_code/yaozhaofang/service/getCaptcha/model"
	"go_code/yaozhaofang/service/getCaptcha/utils"

	"image/color"

	"go_code/yaozhaofang/service/getCaptcha/proto/getCaptcha"
)

type GetCaptcha struct{}

func (e *GetCaptcha) MicroGetCaptcha(ctx context.Context, req *getCaptcha.Request, rsp *getCaptcha.Response) error {
	log.Log("Received GetCaptcha.Call request")
	//	初始化对象
	cap := captcha.New()
	//	设置字体
	cap.SetFont("./conf/comic.ttf")

	//	设置验证码大小
	cap.SetSize(128, 64)
	//	设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	//	设置前景色
	cap.SetFrontColor(color.RGBA{255, 245, 247, 255})
	//	设置背景色
	cap.SetBkgColor(color.RGBA{101, 147, 74, 128}, color.RGBA{69, 137, 148, 255},
		color.RGBA{255, 150, 125, 255})
	//	生成字体 ,将验证码展示到网页

	img, str := cap.Create(4, captcha.NUM)
	//存储图片验证码str到redis
	err := model.SaveImgCode(str, req.Uuid)
	if err != nil{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	//生成的图片序列化
	imgBuf, err := json.Marshal(img)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return err
	}
	//imgBuf用rsp通过消息体定义的img传出
	rsp.Img = imgBuf

	return nil
}
