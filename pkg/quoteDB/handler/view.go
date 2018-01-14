// The handler package contains controllers for displaying web pages.
package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/n7st/quoteDB/pkg/quoteDB/model"

	"github.com/gorilla/mux"
)

// Content{} contains template variables.
type Content struct {
	Error   string
	Title   string
	Channel string
	ID      uint
	Date    time.Time
	Lines   []model.Line
}

// ViewHandler() displays an individual quote.
func (h *Handler) ViewHandler(w http.ResponseWriter, r *http.Request) {
	var lines []model.Line

	vars := mux.Vars(r)
	content := &Content{}

	h.DB.Preload("Head.Channel").Find(&lines, "head_id = ?", vars["id"])

	if len(lines) == 0 {
		content.Title = "No such quote"
		content.Error = "Was an invalid ID used?"
	} else {
		content.Title = fmt.Sprintf("%s at %s", lines[0].Head.Channel.Name, lines[0].CreatedAt)
		content.Lines = lines
		content.Channel = lines[0].Head.Channel.Name
		content.Date = lines[0].Head.CreatedAt.UTC()
		content.ID = lines[0].Head.ID
	}

	err := templates.ExecuteTemplate(w, "quote", content)

	if err != nil {
		log.Println(err)

		fmt.Fprintf(w, "An internal server error occurred: %s", err)
	}
}
