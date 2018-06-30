package core

import (
	"github.com/go-yaml/yaml"
	"github.com/hunterlong/statup/types"
	"io/ioutil"
)

type Config types.Config

func LoadConfig() (*Config, error) {
	var config Config
	file, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &config)
	Configs = &config
	return &config, err
}