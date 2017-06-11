package main

import (
	"github.com/n7st/quoteDB/util"
	"github.com/n7st/quoteDB/pkg/ircbot"
	"github.com/n7st/quoteDB/pkg/ircbot/event"
)

func main() {
	config := util.NewConfig("data/config.yaml")
	bot := util.InitIRC(config)
	db := util.InitDB(config)

	quoteBot := ircbot.NewQuoteBot(bot, db, config)

	event.Initialise(quoteBot)

	quoteBot.IRC.Loop()
}
