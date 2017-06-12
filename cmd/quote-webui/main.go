package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/n7st/quoteDB/util"
	"github.com/n7st/quoteDB/pkg/quote-webui/handler"
)

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
