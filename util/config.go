package util

import (
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"
)

const DEFAULT_CONFIG_FILE_LOCATION = "./data/config.yaml"

type Config struct {
	Channels     []string `yaml:"channels"`
	Debug        bool     `yaml:"debug"`
	UseTLS       bool     `yaml:"use_tls"`
	Verbose      bool     `yaml:"verbose"`
	Port         int      `yaml:"port"`
	AdminChannel string   `yaml:"admin_channel"`
	DBPath       string   `yaml:"db_path"`
	Ident        string   `yaml:"ident"`
	Modes        string   `yaml:"modes"`
	Nickname     string   `yaml:"nickname"`
	Password     string   `yaml:"nickserv_password"`
	Server       string   `yaml:"server"`
	Trigger      string   `yaml:"command_trigger"`
}

func NewConfig(params ...string) *Config {
	var configFileLocation string
	var config = &Config{}

	if len(params) > 0 {
		configFileLocation = params[0]
	} else {
		configFileLocation = DEFAULT_CONFIG_FILE_LOCATION
	}

	data, err := ioutil.ReadFile(configFileLocation)

	if err != nil {
		log.Fatal(err) // We cannot continue without configuration
	}

	err = yaml.Unmarshal(data, config)

	if err != nil {
		log.Fatal(err)
	}

	// Fall back to a standard port
	if config.Port == 0 {
		config.Port = 6667
	}

	if config.Trigger == "" {
		config.Trigger = "!"
	}

	return config

}
