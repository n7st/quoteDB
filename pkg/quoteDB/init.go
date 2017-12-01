package quoteDB

import (
	"github.com/n7st/quoteDB/pkg/quoteDB/util"

	"github.com/jinzhu/gorm"
	"github.com/thoj/go-ircevent"
)

type QuoteBot struct {
	Config *util.Config
	DB     *gorm.DB
	IRC    *irc.Connection

	// Channel -> User -> Message
	History map[string][]map[string]string

	Recover bool
}

// Creates a QuoteBot struct.
func NewQuoteBot(bot *irc.Connection, db *gorm.DB, config *util.Config) *QuoteBot {
	return &QuoteBot{
		Config: config,
		DB:     db,
		IRC:    bot,

		History: make(map[string][]map[string]string),
		Recover: false,
	}
}

// Join a provided list of channels
func (q *QuoteBot) JoinChannels(channels []string) {
	for _, channel := range channels {
		q.IRC.Join(channel)
	}
}
