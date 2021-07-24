module main

go 1.13

replace (
	app => ./src
	model => ./src/model
	scenario => ./src/scenario
)

require (
	app v0.0.0-00010101000000-000000000000
	github.com/RuiHirano/simframe/app v0.0.0-20210724043709-7bcde72837fc // indirect
	github.com/RuiHirano/simframe/engine v0.0.0-20210724043709-7bcde72837fc // indirect
	github.com/fatih/color v1.12.0
	github.com/google/uuid v1.3.0 // indirect
	google.golang.org/grpc v1.39.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	model v0.0.0-00010101000000-000000000000 // indirect
	scenario v0.0.0-00010101000000-000000000000 // indirect
)
