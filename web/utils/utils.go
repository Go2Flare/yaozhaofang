package utils

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

//初始化micro
func InitMicro() client.Client{
//	初始化客户端
	regOpt := registry.Option(func(options *registry.Options){
		options.Addrs = []string{"172.17.0.1:8500"}
	})
	consulRegistry := consul.NewRegistry(regOpt)

	microClient := micro.NewService(
		micro.Registry(consulRegistry),
	)
	return microClient.Client()
}
