package utils

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-plugins/registry/consul"
)

//初始化micro
func InitMicro() client.Client{
//	初始化客户端
	consulReg:=consul.NewRegistry()
	microClient := micro.NewService(
		micro.Registry(consulReg),
	)
	return microClient.Client()
}
