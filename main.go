package main

import (
	"flag"
	"os"

	"github.com/golang/glog"
	"github.com/micro/micro/cmd"
)

func main() {
	flag.Parse()

	defer glog.Flush()
	os.Setenv("MICRO_REGISTRY", "etcd")
	os.Setenv("MICRO_REGISTRY_ADDRESS", "etcd:2379")

	cmd.Init(
		//micro.Name("wpt.api"),
		//micro.Registry(
		//	etcd.NewRegistry(
		//		registry.Addrs("127.0.0.1:2379"),
		//	),
		//),
	)

	//if err := run(); err != nil {
	//	glog.Fatal(err)
	//}
}
