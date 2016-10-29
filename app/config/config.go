package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Listen  string
		LogFile string `yaml:"log_file"`
	}
	Application struct {
		LogFile  string `yaml:"log_file"`
		LogLevel string `yaml:"log_level"`
	}
}

func New(filePath string) (*Config, error) {
	c := &Config{}
	yml, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(yml, &c); err != nil {
		return nil, fmt.Errorf("Error parsing config file: %s", err)
	}
	return c, nil
}
