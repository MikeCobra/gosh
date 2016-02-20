package main

import (
	"io/ioutil"
  "path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Path []string `yaml:",flow"`
}

func LoadConfig() (config Config) {
  filename, _ := filepath.Abs("/home/michaelc/.gosh.yaml")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		panic(err)
	}
  return config
}
