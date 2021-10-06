module github.com/taadis/gateway-service

go 1.14

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/micro v1.18.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
