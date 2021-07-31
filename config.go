package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config represents the base configuration information for the site
type Config struct {
	Title string `yaml:"title"`
}

// ReadConfig loads the configuration data from the `config.yaml` file in the base input directory
func ReadConfig(path string) (*Config, error) {
	var c Config
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(f, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
