package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/consul"
	"go_code/yaozhaofang/service/user/handler"
	"go_code/yaozhaofang/service/user/model"
	user "go_code/yaozhaofang/service/user/proto/user"
)

func main() {
	//初始化配置
	model.InitConfig()
	//初始化redis连接池
	model.InitRedis()

	//初始化服务发现
	regOpt := registry.Option(func(options *registry.Options){
		options.Addrs= []string{"172.17.0.1"}
	})
	consulRegistry := consul.NewRegistry(regOpt)

	// New Service
	service := micro.NewService(
		micro.Address(":52667"),  //固定端口
		micro.Name("go.micro.srv.user"),
		micro.Registry(consulRegistry),   //注册服务
		micro.Version("latest"),
	)
	// Initialise service
	service.Init()
	//在这里用配置文件初始化
	model.InitDb()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	//model.GlobalDB.Close()
}
