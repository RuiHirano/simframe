module simframe-engine

go 1.13

replace (
	app => ./src
	model => ./src/model
	scenario => ./src/scenario
)

require (
	app v0.0.0-00010101000000-000000000000
	github.com/RuiHirano/simframe/app v0.0.0-20210720100252-33c11f9006dc // indirect
	github.com/RuiHirano/simframe/app/config v0.0.0-20210720100252-33c11f9006dc // indirect
	github.com/RuiHirano/simframe/app/model v0.0.0-20210720100252-33c11f9006dc // indirect
	github.com/RuiHirano/simframe/app/scenario v0.0.0-20210720100252-33c11f9006dc // indirect
	github.com/RuiHirano/simframe/engine v0.0.0-20210720100252-33c11f9006dc
	github.com/RuiHirano/simframe/pads/master v0.0.0-20210720100252-33c11f9006dc // indirect
	github.com/RuiHirano/simframe/pads/proto v0.0.0-20210720100252-33c11f9006dc // indirect
	github.com/RuiHirano/simframe/pads/worker v0.0.0-20210720100252-33c11f9006dc // indirect
	github.com/RuiHirano/simframe/util v0.0.0-20210720100252-33c11f9006dc // indirect
	github.com/fatih/color v1.12.0
	google.golang.org/grpc v1.39.0 // indirect
	model v0.0.0-00010101000000-000000000000 // indirect
	scenario v0.0.0-00010101000000-000000000000 // indirect
)
