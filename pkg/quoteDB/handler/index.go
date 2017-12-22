// The handler package contains controllers for displaying web pages.
package handler

import (
	"fmt"
	"log"
	"net/http"
	"github.com/n7st/quoteDB/pkg/quoteDB/model"
)

type IndexContent struct {
	Channels []model.Channel
}

// IndexHandler() displays the website's front page.
func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	content := &IndexContent{}

	h.DB.Find(&content.Channels, model.Channel{})

	err := templates.ExecuteTemplate(w, "index", content)

	if err != nil {
		log.Println(err)

		fmt.Fprintf(w, "An internal server error occurred: %s", err)
	}
}
