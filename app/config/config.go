package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Listen struct {
		Address string
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
