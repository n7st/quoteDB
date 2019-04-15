// quoteDB spawns an IRC frontend to collect quotes from channels and a web UI
// for viewing them.
package main

import (
	"log"
	"net/http"
	"os"

	"git.netsplit.uk/mike/quoteDB/pkg/quoteDB"
	"git.netsplit.uk/mike/quoteDB/pkg/quoteDB/event"
	"git.netsplit.uk/mike/quoteDB/pkg/quoteDB/handler"
	"git.netsplit.uk/mike/quoteDB/pkg/quoteDB/util"

	"github.com/gorilla/handlers"
)

// main() is the program's main loop. It instantiates the bot and web UI and
// connects them to IRC until they are killed.
func main() {
	var config *util.Config

	if len(os.Args) > 1 {
		config = util.NewConfig(os.Args[1])
	} else {
		config = util.NewConfig()
	}

	if config.Server == "" {
		log.Fatal("Invalid configuration provided")
	}

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
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}

	// Run the IRC bot
	quoteBot.IRC.Loop()
}
