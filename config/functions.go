package config

import (
	"encoding/json"
	"log"
	"os"
)

var conf Conf

func Init() Conf {
	configFile, err := os.Open("env/config.json")
	if err != nil {
		log.Fatal("error in reading config file")
	}
	jsonParser := json.NewDecoder(configFile)
	if err != nil {
		log.Fatal("error in marshalling the file")
	}
	err = jsonParser.Decode(&conf)
	if err != nil {
		log.Fatal("error in unmarshalling file")
	}
	return conf
}
