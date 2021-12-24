package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Pets []Pet
}

var configuration Config

type Pet struct {
	Name string `yaml:"name"`
	Form string `yaml:"type"`
}

func NewConfig() (c *Config) {
	c = new(Config)
	return
}

func LoadFromFile(filename string) (c *Config, err error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to load config file [%v],error=%v", filename, err)
	}
	pet := NewConfig()
	err = yaml.Unmarshal(content, &pet)
	if err != nil {
		return nil, fmt.Errorf("failed to parse yaml data,error=%v", err)
	}
	return pet, err
}
func Set(c *Config) {
	configuration = *c
}
func Get() (c *Config) {
	return &configuration
}
