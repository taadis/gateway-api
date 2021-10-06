#export MICRO_REGISTRY=etcd
#export MICRO_REGISTRY_ADDRESS=127.0.0.1:2379
#export MICRO_API_HANDLER=http

#.PHONY: build
#build:
#	go build -o mymicro main.go

#.PHONY: runregistry
#runregistry:
#	go run main.go registry

.PHONY: runapi
runapi:
	go run main.go api #--enable_rpc --namespace go.micro

.PHONY: runweb
runweb:
	go run main.go web

