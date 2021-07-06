module github.com/RuiHirano/simframe/src

replace (
	github.com/RuiHirano/simframe/engine => ../engine
	github.com/RuiHirano/simframe/model => ../model
	github.com/RuiHirano/simframe/scenario => ../scenario
	github.com/RuiHirano/simframe/src/model => ./model
	github.com/RuiHirano/simframe/util => ../util
)

go 1.13

require (
	github.com/RuiHirano/simframe/engine v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/simframe/model v0.0.0-00010101000000-000000000000
	github.com/RuiHirano/simframe/scenario v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/simframe/util v0.0.0-00010101000000-000000000000
)
