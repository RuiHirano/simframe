module github.com/RuiHirano/simframe/src

go 1.13

replace (
	github.com/RuiHirano/simframe/app => ../../../../app
	github.com/RuiHirano/simframe/app/config => ../../../../app/config
	github.com/RuiHirano/simframe/app/model => ../../../../app/model
	github.com/RuiHirano/simframe/app/scenario => ../../../../app/scenario
	github.com/RuiHirano/simframe/engine => ../../../../engine
	github.com/RuiHirano/simframe/src/scenario => ./scenario
)

require (
	github.com/RuiHirano/simframe/app v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/simframe/app/config v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/simframe/app/model v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/simframe/app/scenario v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/simframe/builder v0.0.0-20210719093920-5716d9a7fc2f // indirect
	github.com/RuiHirano/simframe/config v0.0.0-20210711094948-052f1fdaf5f6 // indirect
	github.com/RuiHirano/simframe/engine v0.0.0-20210711094948-052f1fdaf5f6 // indirect
	github.com/RuiHirano/simframe/model v0.0.0-20210711094948-052f1fdaf5f6
	github.com/RuiHirano/simframe/scenario v0.0.0-20210711094948-052f1fdaf5f6 // indirect
	github.com/RuiHirano/simframe/util v0.0.0-20210711094948-052f1fdaf5f6
)
