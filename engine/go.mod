module github.com/RuiHirano/simframe/engine

go 1.13

replace (
	github.com/RuiHirano/simframe/api => ../api
	github.com/RuiHirano/simframe/app => ../app
)

require (
	github.com/RuiHirano/simframe/api v0.0.0-20210722053619-b424052e6aad
	github.com/RuiHirano/simframe/app v0.0.0-20210722053619-b424052e6aad
	github.com/google/uuid v1.1.2
	google.golang.org/genproto v0.0.0-20210722135532-667f2b7c528f // indirect
	google.golang.org/grpc v1.39.0
	gopkg.in/yaml.v2 v2.2.3
)
