package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"go_code/yaozhaofang/service/getArea/handler"
	"go_code/yaozhaofang/service/getArea/model"
	"log"

	getArea "go_code/yaozhaofang/service/getArea/proto/getArea"
)

func main() {
	err := model.InitDb()
	if err != nil {
		log.Fatalf("model.InitDb() err = %v", err)
	}
	model.InitRedis()
	//配置注册consul的地址
	regOpt := registry.Option(func(options *registry.Options){
			options.Addrs= []string{"172.17.0.1"}
	})
	consulRegistry := consul.NewRegistry(regOpt)


	// New Service 监听的地址
	service := micro.NewService(
		micro.Address(":52668"),
		micro.Name("go.micro.srv.getArea"),
		micro.Registry(consulRegistry),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getArea.RegisterGetAreaHandler(service.Server(), new(handler.GetArea))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
