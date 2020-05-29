package main

import (
	"io/ioutil"
	"log"
	"path"

	yaml "gopkg.in/yaml.v2"
)

const ConfigFile = "config.yaml"
const BuildVersion = "BUILD_VERSION"

type Config struct {
	Service struct {
		OutputFormat string `yaml:"output_format"`
	} `yaml:"service"`
}

var AppConfig *Config

func LoadConfig(dir string) {
	data, err := ioutil.ReadFile(path.Join(dir, ConfigFile))
	if err != nil {
		log.Panicln(err)
	}
	AppConfig = &Config{}
	err = yaml.Unmarshal(data, AppConfig)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
