package global

import (
	"gopkg.in/yaml.v2"
	"lixiao189/games-admin-server/log"
	"os"
)

type configType struct {
	Serv struct {
		Addr string
	}
	DB struct {
		Hostname string `yaml:"name"`
		Port     string
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	}
}

var Config configType

func initConfig() {
	configFile, err := os.ReadFile(*ConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(configFile, &Config)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(Config)
}
