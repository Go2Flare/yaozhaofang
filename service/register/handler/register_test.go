package handler

import (
	"context"
	"fmt"
	"go_code/yaozhaofang/service/register/model"
	"go_code/yaozhaofang/service/register/proto/register"
	"log"
	"testing"
)

func init(){
	model.InitRedis()
	err := model.InitDb()
	if err != nil {
		log.Fatalf("model.InitDb : %v", err)
	}
}

func TestRegister(t *testing.T){
	r := Register{}
	req := &register.RegRequest{
		Mobile: "13790887214",
		Password: "123456",
		SmsCode: "768475",
	}
	rsp := &register.RegResponse{}
	err := r.Register(context.Background(), req, rsp)
	if err != nil {
		log.Fatalf("r.Register err : %v",err)
	}
	fmt.Println(rsp)
}

func TestSmsCode(t *testing.T){
	r := Register{}
	req := &register.Request{
		Mobile: "13790887214",
		Text: "123", //前端传来的验证码
		Uuid: "2222", //redis中的验证码
	}
	rsp := &register.Response{}
	err := r.SmsCode(context.Background(), req, rsp)
	if err != nil{
		log.Fatalf("r.SmsCode : %v", err)
	}
}
