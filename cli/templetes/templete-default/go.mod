module github.com/RuiHirano/simframe/src

go 1.13

replace (
	app => ./src
	github.com/RuiHirano/simframe/app => ../../../app
	github.com/RuiHirano/simframe/app/config => ../../../app/config
	github.com/RuiHirano/simframe/app/model => ../../../app/model
	github.com/RuiHirano/simframe/app/scenario => ../../../app/scenario
	github.com/RuiHirano/simframe/builder => ../../../builder
	github.com/RuiHirano/simframe/engine => ../../../engine
	github.com/RuiHirano/simframe/pads => ../../../pads
	github.com/RuiHirano/simframe/pads/master => ../../../pads/master
	github.com/RuiHirano/simframe/pads/proto => ../../../pads/proto
	github.com/RuiHirano/simframe/pads/worker => ../../../pads/worker
	github.com/RuiHirano/simframe/util => ../../../util
	model => ./src/model
	scenario => ./src/scenario
)

require (
	app v0.0.0-00010101000000-000000000000
	github.com/RuiHirano/simframe/app v0.0.0-00010101000000-000000000000
	github.com/RuiHirano/simframe/app/config v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/simframe/app/scenario v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/simframe/builder v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/simframe/engine v0.0.0-00010101000000-000000000000
	github.com/RuiHirano/simframe/pads/master v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/simframe/pads/proto v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/simframe/pads/worker v0.0.0-00010101000000-000000000000 // indirect
	github.com/fatih/color v1.7.0
	google.golang.org/grpc v1.39.0 // indirect
	model v0.0.0-00010101000000-000000000000 // indirect
	scenario v0.0.0-00010101000000-000000000000 // indirect
)
