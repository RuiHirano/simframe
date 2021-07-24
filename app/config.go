package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

type IConfig interface {
	GetName() string
	GetConfig() string
}

type Config struct{
	Name string `json:"name"`
	Version string `json:"version"`
}

func NewConfig() *Config {

	conf := &Config{}
	conf.LoadConfig()

	return conf
}

func (cf *Config) LoadConfig(){
	p, _ := os.Getwd()
    currentDirPath := p
	raw, err := ioutil.ReadFile(fmt.Sprintf("%s/simframe.config.json", currentDirPath))
	if err != nil{
		color.Red("cannot find simframe.config.json\n")
		os.Exit(1)
	}
	err = json.Unmarshal(raw, cf)
	if cf == nil || cf.Name == "" || cf.Version == ""{
		color.Red("invalid args in simframe.config.json\n")
		color.Red("%v", raw, fmt.Sprintf("%s/simframe.config.json", currentDirPath), err)
		os.Exit(1)
	}
}

func (cf *Config) GetName() string{
	return cf.Name
}

func (cf *Config) GetVersion() string{
	return cf.Version
}