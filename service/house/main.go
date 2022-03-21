package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"go_code/yaozhaofang/service/house/handler"
	"go_code/yaozhaofang/service/house/model"
	"log"

	house "go_code/yaozhaofang/service/house/proto/house"
)

func main() {
	model.InitDb()

	regOpt := registry.Option(func(options *registry.Options){
		options.Addrs= []string{"172.17.0.1"}
	})
	consulRegistry := consul.NewRegistry(regOpt)

	// New Service
	service := micro.NewService(
		micro.Address(":52671"),
		micro.Name("go.micro.srv.house"),
		micro.Registry(consulRegistry),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	house.RegisterHouseHandler(service.Server(), new(handler.House))
	//
	//// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.house", service.Server(), new(subscriber.House))
	//
	//// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.house", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
