module github.com/RuiHirano/simframe/engine

go 1.13

replace (
	github.com/RuiHirano/simframe/engine/pads/master => ../pads/master
	github.com/RuiHirano/simframe/app => ../app
	github.com/RuiHirano/simframe/builder => ../builder
	github.com/RuiHirano/simframe/engine/pads/worker => ../pads/worker
)

require (
	github.com/RuiHirano/simframe/config v0.0.0-20210711094948-052f1fdaf5f6 // indirect
	github.com/RuiHirano/simframe/model v0.0.0-20210711094948-052f1fdaf5f6
	github.com/RuiHirano/simframe/scenario v0.0.0-20210711094948-052f1fdaf5f6
	github.com/RuiHirano/simframe/util v0.0.0-20210711094948-052f1fdaf5f6 // indirect
	github.com/golang/protobuf v1.5.0
	google.golang.org/grpc v1.39.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
)
