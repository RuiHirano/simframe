package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"path"
	"runtime"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

var (
    templeteDirPath string
    currentDirPath string
)

func init(){
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        panic("No caller information")
    }
    templeteDirPath = path.Dir(path.Dir(filename) + "/../templetes/")
    p, _ := os.Getwd()
    currentDirPath = p
}

func GetProjectName() string{
    for {
        prompt := promptui.Prompt{
            Label:    "What is your project name?",
            Validate: func(input string) error {
                if len(input) < 1 {
                    return errors.New("Project name must have more than 1 characters")
                }
                return nil
            },
            Default:  "my-simframe-project",
        }
        projectName, err := prompt.Run()
        if err != nil {
            color.Red("Prompt failed %v\n", err)
            os.Exit(1)
        }else{
            // Error if project already exist
            if _, err := os.Stat(fmt.Sprintf("%s/%s", currentDirPath, projectName)); err == nil {
                color.Red("Error: Project %s is already exist. %v\n", projectName, err)
                os.Exit(1)
            }
            return projectName
        }
    }
}

func GetTempleteList() []string{
    files, err := ioutil.ReadDir(templeteDirPath)
    if err != nil {
        log.Fatal(err)
    }
 
    tempList := []string{}
    for _, f := range files {
        if f.IsDir(){
            fmt.Println(f.Name())
            tempList = append(tempList, f.Name())
        }
    }
    return tempList
}

func GetTempleteName() string{
    tempList := GetTempleteList()

    for {
        prompt := promptui.Select{
            Label: "Which template do you use?",
            Items: tempList,
        }
        _, templeteName, err := prompt.Run()
        if err != nil {
            color.Red("Prompt failed %v\n", err)
            os.Exit(1)
        }else{
            return templeteName
        }
    }
}

func InstallTemplete(projectName string, templeteName string){
    err := copy.Copy(fmt.Sprintf("%s/%s", templeteDirPath, templeteName), fmt.Sprintf("%s/%s", currentDirPath, projectName))
    if err != nil {
        color.Red("Install templete error %v\n", err)
        os.Exit(1)
    }
}

func CreateProject(projectName string, templeteName string){
    if err := os.Mkdir(projectName, 0777); err != nil {
        fmt.Println(err)
    }

    // install templete code
    InstallTemplete(projectName, templeteName)

    // go build

    color.Green("Project %s is created!\n\n", projectName)
    color.Green("You can run project by below command\n\n")
    color.Green("$ cd %s\n$ simcli start\n", projectName)
}

func Init(){
    projectName := GetProjectName()
    templeteName := GetTempleteName()
	fmt.Printf("Project Name is %q\n", projectName)
	fmt.Printf("Templete Name is %q\n", templeteName)
    CreateProject(projectName, templeteName)
}

var initCmd = &cobra.Command{
    Use:   "init",
    Short: "command line calculator",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Init command")
        Init()
    },
}

func init() {
	RootCmd.AddCommand(initCmd)
}