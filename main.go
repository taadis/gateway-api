package main

import (
	"flag"
	"os"

	_ "github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3/cmd"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()

	defer glog.Flush()
	os.Setenv("MICRO_REGISTRY", "etcd")
	os.Setenv("MICRO_REGISTRY_ADDRESS", "etcd:2379")

	err := cmd.Init()
	if err != nil {
		glog.Fatalf("cmd.Init error:%+v", err)
	}
	//cmd.Init(
	//	micro.Name("wpt.api"),
	//	micro.Registry(
	//		etcd.NewRegistry(
	//			registry.Addrs("127.0.0.1:2379"),
	//		),
	//	),
	//)

	//if err := run(); err != nil {
	//	glog.Fatal(err)
	//}
}
