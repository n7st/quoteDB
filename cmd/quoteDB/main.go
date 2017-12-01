// quoteDB spawns an IRC frontend to collect quotes from channels and a web UI
// for viewing them.
package main

import (
	"net/http"

	"github.com/n7st/quoteDB/pkg/quoteDB"
	"github.com/n7st/quoteDB/pkg/quoteDB/event"
	"github.com/n7st/quoteDB/pkg/quoteDB/util"
	"github.com/n7st/quoteDB/pkg/quoteDB/handler"

	"github.com/gorilla/handlers"
)

// main() is the program's main loop. It instantiates the bot and web UI and
// connects them to IRC until they are killed.
func main() {
	config := util.NewConfig("data/config.yaml")
	bot := util.InitIRC(config)
	db := util.InitDB(config)
	router := handler.NewHandler(db, config).Router()

	quoteBot := quoteDB.NewQuoteBot(bot, db, config)

	srv := &http.Server{
		Handler: handlers.CORS()(router),
		Addr:    ":" + config.WebUIPort,
	}

	defer db.Close()

	event.Initialise(quoteBot)

	// Run the web server
	srv.ListenAndServe()

	// Run the IRC bot
	quoteBot.IRC.Loop()
}
