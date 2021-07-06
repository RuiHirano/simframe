module github.com/RuiHirano/simframe/src

replace (
 github.com/RuiHirano/simframe/model => ../model
github.com/RuiHirano/simframe/engine => ../engine
github.com/RuiHirano/simframe/src/model => ./model
)

go 1.13

require (
	github.com/RuiHirano/simframe/engine v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/simframe/model v0.0.0-00010101000000-000000000000
)
