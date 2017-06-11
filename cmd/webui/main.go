package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/n7st/quoteDB/util"
	"github.com/n7st/quoteDB/pkg/webui/handler"
	"log"
)

func main() {
	config := util.NewConfig("data/config.yaml")
	db := util.InitDB(config)
	router := handler.NewHandler(db).Router()

	defer db.Close()

	srv := &http.Server{
		Handler: handlers.CORS()(router),
		Addr:    ":8080",
	}

	log.Fatal(srv.ListenAndServe())
}
