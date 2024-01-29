package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type (
	AppConfig struct {
		data string `yaml: "data"`
	}
	ConfigFile map[string]*AppConfig
)

func LoadConfig(env string) (*AppConfig, error) {

}

func main() {

}
