module github.com/RuiHirano/simframe/cli

replace github.com/RuiHirano/simframe/cli/cmd => ./cmd

go 1.13

require (
	github.com/c-bata/go-prompt v0.2.6
	github.com/fatih/color v1.7.0
	github.com/gdamore/tcell/v2 v2.3.3
	github.com/manifoldco/promptui v0.8.0
	github.com/otiai10/copy v1.6.0
	github.com/rivo/tview v0.0.0-20210624165335-29d673af0ce2
	github.com/spf13/cobra v1.2.1
)
