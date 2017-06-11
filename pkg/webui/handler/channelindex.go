package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/n7st/quoteDB/model"
)

type IndexContent struct {
	Instances []model.Head
	Channel   string
	Error     string
}

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	var head []model.Head

	vars := mux.Vars(r)
	content := &IndexContent{Channel: vars["name"]}

	h.DB.Find(&head, model.Head{Channel: vars["name"]})

	if len(head) == 0 {
		content.Error = "Could not find any quotes"
	} else {
		content.Instances = head
	}

	err := templates.ExecuteTemplate(w, "channel_index", content)

	if err != nil {
		log.Println(err)
	}
}
