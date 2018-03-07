package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Websites []WebSite `yaml:"websites"`
	Port     string    `yaml:"port"`
	DBFile   string    `yaml:"db_file"`
}

type WebSite struct {
	URLFilter        string `yaml:"url_filter"`
	DateMatch        string `yaml:"date_match"`
	DateLayout       string `yaml:"date_layout"`
	TimeZone         string `yaml:"timezone"`
	ReturnUnmodified bool   `yaml:"return_unmodified"`
}

func NewConfig(yamlpath string) (*Config, error) {
	bin, err := ioutil.ReadFile(yamlpath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(bin, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
