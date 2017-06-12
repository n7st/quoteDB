// The util package provides functionality needed by all binaries in the
// project.
package util

import (
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"
)

// DEFAULT_CONFIG_FILE_LOCATION can be overridden by passing a command-line
// argument to the program.
const DEFAULT_CONFIG_FILE_LOCATION = "./data/config.yaml"

// Config{} contains configuration values for the bot and web server.
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
	BaseURL      string   `yaml:"base_url"`
	WebUIPort    string   `yaml:"webui_port"`
}

// NewConfig() creates a new Config{} from a YAML file and sets some defaults.
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
