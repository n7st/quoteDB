// The handler package contains controllers for displaying web pages.
package handler

import (
	"fmt"
	"log"
	"net/http"
)

// IndexHandler() displays the website's front page.
func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index", "")

	if err != nil {
		log.Println(err)

		fmt.Fprintf(w, "An internal server error occurred: %s", err)
	}
}
