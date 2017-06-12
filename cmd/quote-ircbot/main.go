package main

import (
	"github.com/n7st/quoteDB/util"
	"github.com/n7st/quoteDB/pkg/quote-ircbot"
	"github.com/n7st/quoteDB/pkg/quote-ircbot/event"
)

func main() {
	config := util.NewConfig("data/config.yaml")
	bot := util.InitIRC(config)
	db := util.InitDB(config)

	quoteBot := quote_ircbot.NewQuoteBot(bot, db, config)

	defer db.Close()

	event.Initialise(quoteBot)

	quoteBot.IRC.Loop()
}
