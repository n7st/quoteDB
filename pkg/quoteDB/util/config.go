// The util package provides functionality needed by all binaries in the
// project.
package util

import (
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"
	"github.com/shibukawa/configdir"
)

const (
	configFileName  = "config.yaml"
	vendorName      = "netsplit"
	applicationName = "quoteDB"

	defaultPort           = 6667
	defaultTrigger        = "!"
	defaultMaxHistorySize = 200
	defaultMaxQuoteSize   = 50
)

// Config{} contains configuration values for the bot and web server.
type Config struct {
	Channels       []string `yaml:"channels"`
	Debug          bool     `yaml:"debug"`
	UseTLS         bool     `yaml:"use_tls"`
	Verbose        bool     `yaml:"verbose"`
	Port           int      `yaml:"port"`
	MaxQuoteSize   int      `yaml:"max_quote_size"`
	MaxHistorySize int      `yaml:"max_history_size"`
	AdminChannel   string   `yaml:"admin_channel"`
	DBPath         string   `yaml:"db_path"`
	Ident          string   `yaml:"ident"`
	Modes          string   `yaml:"modes"`
	Nickname       string   `yaml:"nickname"`
	Password       string   `yaml:"nickserv_password"`
	Server         string   `yaml:"server"`
	Trigger        string   `yaml:"command_trigger"`
	BaseURL        string   `yaml:"base_url"`
	WebUIPort      string   `yaml:"webui_port"`
}

// NewConfig() creates a new Config{} from a YAML file and sets some defaults.
func NewConfig(params ...string) *Config {
	var configFileLocation string
	var config = &Config{}
	var data []byte
	var err error

	if len(params) > 0 {
		configFileLocation = params[0]

		data, err = ioutil.ReadFile(configFileLocation)
	} else {
		configDirs := configdir.New(vendorName, applicationName)
		folder := configDirs.QueryFolderContainsFile(configFileName)

		if folder != nil {
			log.Println("Loading config from default location")

			data, err = folder.ReadFile(configFileName)
		}
	}

	if err != nil {
		log.Fatal(err) // We cannot continue without configuration
	}

	err = yaml.Unmarshal(data, config)

	if err != nil {
		log.Fatal(err)
	}

	// Fall back to a standard port
	if config.Port == 0 {
		config.Port = defaultPort
	}

	if config.Trigger == "" {
		config.Trigger = defaultTrigger
	}

	if config.MaxQuoteSize == 0 {
		config.MaxQuoteSize = defaultMaxQuoteSize
	}

	if config.MaxHistorySize == 0 {
		config.MaxHistorySize = defaultMaxHistorySize
	}

	return config
}
