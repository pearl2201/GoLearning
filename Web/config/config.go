package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	ServerPort string
}

var err error
var config Configuration

func ReadConfig(filename string) (Configuration, error) {
	configFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Unable to read config, swtiching to flag mode")
		return Configuration{}, err
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Println("Invalid json, expecting port from command line flag")
		return Configuration{}, err
	}

	return config, nil
}
