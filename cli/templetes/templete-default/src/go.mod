module github.com/RuiHirano/simframe/src

go 1.13

replace (
	github.com/RuiHirano/simframe/src/scenario => ./scenario
)

require (
	github.com/RuiHirano/simframe/app/config v0.0.0-20210711094948-052f1fdaf5f6 // indirect
	github.com/RuiHirano/simframe/engine v0.0.0-20210711094948-052f1fdaf5f6 // indirect
	github.com/RuiHirano/simframe/model v0.0.0-20210711094948-052f1fdaf5f6
	github.com/RuiHirano/simframe/scenario v0.0.0-20210711094948-052f1fdaf5f6
	github.com/RuiHirano/simframe/util v0.0.0-20210711094948-052f1fdaf5f6
)
