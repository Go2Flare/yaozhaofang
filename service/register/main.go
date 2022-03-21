package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/consul"
	//"github.com/micro/protoc-gen-micro/plugin/micro"
	"go_code/yaozhaofang/service/register/handler"
	"go_code/yaozhaofang/service/register/model"
	register "go_code/yaozhaofang/service/register/proto/register"
)

func main() {
	//初始化数据库配置
	model.InitConfig()
	//初始化mysql,redis
	model.InitRedis()
	model.InitDb()

	//服务发现consul
	regOpt := registry.Option(func(options *registry.Options){
		options.Addrs= []string{"172.17.0.1"}
	})
	consulRegistry := consul.NewRegistry(regOpt)

	// New Service
	service := micro.NewService(
		micro.Address(":52669"),
		micro.Name("go.micro.srv.register"),
		micro.Version("lastest"),
		micro.Registry(consulRegistry),
	)

	// Initialise service
	service.Init()

	// Register Handler
	register.RegisterRegisterHandler(service.Server(), new(handler.Register))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
