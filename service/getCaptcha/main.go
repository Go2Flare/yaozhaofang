package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/consul"
	//"github.com/micro/protoc-gen-micro/plugin/micro"
	"go_code/yaozhaofang/service/getCaptcha/handler"
	"go_code/yaozhaofang/service/getCaptcha/model"
	getCaptcha "go_code/yaozhaofang/service/getCaptcha/proto/getCaptcha"
)

func main() {
	//初始化consul注册服务
	regOpt := registry.Option(func(options *registry.Options){
		options.Addrs= []string{"172.17.0.1"}
	})
	consulRegistry := consul.NewRegistry(regOpt)

	//初始化reds
	model.InitRedis()

	// New Service
	service := micro.NewService(
		//指定微服务连接服务发现的端口
		micro.Address(":52666"), //防止随机生成port
		micro.Name("go.micro.srv.getCaptcha"),
		micro.Registry(consulRegistry),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getCaptcha.RegisterGetCaptchaHandler(service.Server(), new(handler.GetCaptcha))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
