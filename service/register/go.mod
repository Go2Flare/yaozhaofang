module go_code/yaozhaofang/service/register

go 1.14

require (
	github.com/alibabacloud-go/darabonba-openapi v0.1.12
	github.com/alibabacloud-go/dysmsapi-20170525/v2 v2.0.8
	github.com/alibabacloud-go/tea v1.1.17
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190808125512-07798873deee // indirect
	github.com/gin-contrib/sessions v0.0.4
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/spf13/viper v1.8.1
	google.golang.org/grpc/examples v0.0.0-20210915223801-7cf9689be2d2 // indirect
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => google.golang.org/grpc v1.25.1
