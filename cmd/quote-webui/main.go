// quote-webui is the web frontend for this quote database software and it is
// used for displaying quotes on web pages in a pretty format.
package main

import (
	"log"
	"net/http"

	"github.com/n7st/quoteDB/util"
	"github.com/n7st/quoteDB/pkg/quote-webui/handler"

	"github.com/gorilla/handlers"
)

// main() is the program's main loop. It launches the web server and runs it
// until it is killed.
func main() {
	config := util.NewConfig("data/config.yaml")
	db := util.InitDB(config)
	router := handler.NewHandler(db, config).Router()

	defer db.Close()

	srv := &http.Server{
		Handler: handlers.CORS()(router),
		Addr:    ":"+config.WebUIPort,
	}

	log.Fatal(srv.ListenAndServe())
}
