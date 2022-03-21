package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/consul"
	//"github.com/micro/protoc-gen-micro/plugin/micro"
	"go_code/yaozhaofang/service/userOrder/handler"
	"go_code/yaozhaofang/service/userOrder/model"
	userOrder "go_code/yaozhaofang/service/userOrder/proto/userOrder"
)

func main() {
	model.InitConfig()
	model.InitDb()

	regOpt := registry.Option(func(options *registry.Options){
		options.Addrs= []string{"172.17.0.1"}
	})
	consulRegistry := consul.NewRegistry(regOpt)

	// New Service
	service := micro.NewService(
		micro.Address(":52670"),
		micro.Name("go.micro.srv.userOrder"),
		micro.Registry(consulRegistry),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	userOrder.RegisterUserOrderHandler(service.Server(), new(handler.UserOrder))
	//
	//// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.userOrder", service.Server(), new(subscriber.UserOrder))
	//
	//// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.userOrder", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
