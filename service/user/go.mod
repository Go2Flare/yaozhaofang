module go_code/yaozhaofang/service/user

go 1.14

require (
	github.com/gin-contrib/sessions v0.0.4
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/spf13/viper v1.8.1
	github.com/tedcy/fdfs_client v0.0.0-20200106031142-21a04994525a
	google.golang.org/grpc/examples v0.0.0-20210914231103-2d4e44a0cd75 // indirect
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
