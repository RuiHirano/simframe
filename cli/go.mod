module github.com/RuiHirano/simframe/cli

go 1.13

replace (
	github.com/RuiHirano/simframe/app => ../app
	github.com/RuiHirano/simframe/app/config => ../app/config
	github.com/RuiHirano/simframe/app/scenario => ../app/scenario
	github.com/RuiHirano/simframe/builder => ../builder
)

require (
	github.com/RuiHirano/simframe/app v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/simframe/app/config v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/simframe/app/scenario v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/simframe/builder v0.0.0-00010101000000-000000000000
	github.com/briandowns/spinner v1.16.0
	github.com/fatih/color v1.12.0
	github.com/manifoldco/promptui v0.8.0
	github.com/otiai10/copy v1.6.0
	github.com/spf13/cobra v1.2.1
)
