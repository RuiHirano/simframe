module github.com/RuiHirano/simframe/engine

go 1.13

replace (
	github.com/RuiHirano/simframe/app => ../app
	github.com/RuiHirano/simframe/builder => ../builder
	github.com/RuiHirano/simframe/engine/api => ../engine/proto
	github.com/RuiHirano/simframe/engine/master => ../engine/master
	github.com/RuiHirano/simframe/engine/worker => ../engine/worker
)

require (
	github.com/RuiHirano/simframe/app v0.0.0-00010101000000-000000000000
	github.com/RuiHirano/simframe/app/config v0.0.0-20210721053332-f8df2845984a // indirect
	github.com/RuiHirano/simframe/app/model v0.0.0-20210721053332-f8df2845984a // indirect
	github.com/RuiHirano/simframe/app/scenario v0.0.0-20210721053332-f8df2845984a // indirect
	github.com/RuiHirano/simframe/engine/api v0.0.0-00010101000000-000000000000
	github.com/RuiHirano/simframe/engine/master v0.0.0-00010101000000-000000000000
	github.com/RuiHirano/simframe/engine/worker v0.0.0-00010101000000-000000000000
	github.com/RuiHirano/simframe/pads/proto v0.0.0-20210721041131-f9025356f0f2 // indirect
	github.com/RuiHirano/simframe/pads/worker v0.0.0-20210721041131-f9025356f0f2 // indirect
	github.com/RuiHirano/simframe/util v0.0.0-20210721053332-f8df2845984a // indirect
	google.golang.org/grpc v1.39.0
)
