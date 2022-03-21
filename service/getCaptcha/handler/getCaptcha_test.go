package handler

import (
	"context"
	"fmt"
	"go_code/yaozhaofang/service/getCaptcha/model"
	"go_code/yaozhaofang/service/getCaptcha/proto/getCaptcha"
	"testing"
)

func init() {
	model.InitRedis()
}

func TestMicroGetCaptcha(t *testing.T) {
	g := GetCaptcha{}
	rsp := &getCaptcha.Response{}
	g.MicroGetCaptcha(context.Background(), &getCaptcha.Request{
		Uuid: "123456",
	}, rsp)
	fmt.Println(rsp)
}
