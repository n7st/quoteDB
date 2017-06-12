// quote-ircbot is an IRC frontend for this quote database software. It is
// capable of grabbing multiple lines at once and storing them, and intended to
// be used in tandem with quote-webui for displaying quotes.
package main

import (
	"github.com/n7st/quoteDB/pkg/quote-ircbot"
	"github.com/n7st/quoteDB/pkg/quote-ircbot/event"
	"github.com/n7st/quoteDB/util"
)

// main() is the program's main loop. It instantiates the bot and connects it to
// IRC until it is killed.
func main() {
	config := util.NewConfig("data/config.yaml")
	bot := util.InitIRC(config)
	db := util.InitDB(config)

	quoteBot := quote_ircbot.NewQuoteBot(bot, db, config)

	defer db.Close()

	event.Initialise(quoteBot)

	quoteBot.IRC.Loop()
}
