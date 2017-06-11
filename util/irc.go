package util

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/thoj/go-ircevent"
)

func InitIRC(config *Config) *irc.Connection {
	bot := irc.IRC(config.Nickname, config.Ident)

	bot.VerboseCallbackHandler = config.Verbose
	bot.Debug = config.Debug
	bot.UseTLS = config.UseTLS
	bot.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	server := fmt.Sprintf("%s:%d", config.Server, config.Port)

	err := bot.Connect(server)

	if err != nil {
		log.Fatal(err)
	}

	return bot
}
