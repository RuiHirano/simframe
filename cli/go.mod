module github.com/RuiHirano/simframe/cli

replace github.com/RuiHirano/simframe/cli/cmd => ./cmd

replace github.com/RuiHirano/simframe/builder => ../builder

go 1.13

require (
	github.com/RuiHirano/simframe/builder v0.0.0-00010101000000-000000000000
	github.com/briandowns/spinner v1.16.0
	github.com/fatih/color v1.12.0
	github.com/manifoldco/promptui v0.8.0
	github.com/otiai10/copy v1.6.0
	github.com/spf13/cobra v1.2.1
	gopkg.in/yaml.v2 v2.4.0
)
