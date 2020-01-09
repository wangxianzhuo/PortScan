package config

import (
	"encoding/json"
	"io/ioutil"
)

// Load load configurations
func Load(file string) (Configuration, error) {
	var conf Configuration
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return Configuration{}, err
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return Configuration{}, err
	}

	return conf, nil
}
