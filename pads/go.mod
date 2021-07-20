module github.com/RuiHirano/simframe/pads

go 1.13

replace (
    github.com/RuiHirano/simframe/app/config => ../app/config
    github.com/RuiHirano/simframe/app/scenario => ../app/scenario
    github.com/RuiHirano/simframe/pads/master => ./master
    github.com/RuiHirano/simframe/pads/worker => ./worker
)