package apiserver

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Config ...
type Config struct {
	ServerAddress string
	ServerPort    string
}

// GetConfig ...
func (config *Config) GetConfig() *Config {
	yamlText, err := ioutil.ReadFile("/server/internal/config/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlText, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return config
}
