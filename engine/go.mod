module github.com/RuiHirano/simframe/engine

go 1.13


replace (
    github.com/RuiHirano/simframe/app => ../app
    github.com/RuiHirano/simframe/builder => ../builder
    github.com/RuiHirano/simframe/pads/master => ../pads/master
    github.com/RuiHirano/simframe/pads/worker => ../pads/worker
)

require (
)
