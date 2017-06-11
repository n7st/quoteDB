package ircbot

import (
	"github.com/thoj/go-ircevent"
	"github.com/jinzhu/gorm"

	"github.com/n7st/quoteDB/util"
)

type QuoteBot struct {
	Config *util.Config
	DB     *gorm.DB
	IRC    *irc.Connection

	// Channel -> User -> Message
	History map[string][]map[string]string
}

func NewQuoteBot(bot *irc.Connection, db *gorm.DB, config *util.Config) *QuoteBot {
	return &QuoteBot{
		Config: config,
		DB:     db,
		IRC:    bot,

		History: make(map[string][]map[string]string),
	}
}

func (q *QuoteBot) JoinChannels(channels []string) {
	if len(channels) == 0 {
		return
	}

	for _, channel := range channels {
		q.IRC.Join(channel)
	}
}
