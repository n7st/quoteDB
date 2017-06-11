package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/n7st/quoteDB/model"
)

type Content struct {
	Error   string
	Title   string
	Channel string
	ID      uint
	Date    time.Time
	Lines   []model.Line
}

var templates = template.Must(template.ParseGlob("view/*"))

func (h *Handler) ViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lines := []model.Line{}
	content := &Content{}

	h.DB.Preload("Head").Find(&lines, "head_id = ?", vars["id"])

	if len(lines) == 0 {
		content.Title = "No such quote"
		content.Error = "Was an invalid ID used?"
	} else {
		content.Title = fmt.Sprintf("%s at %s", lines[0].Head.Channel, lines[0].CreatedAt)
		content.Lines = lines
		content.Channel = lines[0].Head.Channel
		content.Date = lines[0].Head.CreatedAt.UTC()
		content.ID = lines[0].Head.ID
	}

	err := templates.ExecuteTemplate(w, "quote", content)

	if err != nil {
		log.Fatal(err)
	}
}
