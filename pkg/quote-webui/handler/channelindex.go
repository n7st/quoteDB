// The handler package contains controllers for displaying web pages.
package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/n7st/quoteDB/model"

	"github.com/gorilla/mux"
)

// ChannelIndexContent{} contains template variables.
type ChannelIndexContent struct {
	Instances []model.Head
	Channel   string
	Error     string
	Trigger   string
}

// headWithLines{} stores a full quote.
type headWithLines struct {
	Head  model.Head
	Lines []model.Line
}

// ChannelIndexHandler() displays a page with a list of quotes for the given
// {channel}.
func (h *Handler) ChannelIndexHandler(w http.ResponseWriter, r *http.Request) {
	var (
		head []model.Head
		out  []model.Head
	)

	vars := mux.Vars(r)
	content := &ChannelIndexContent{
		Channel: vars["name"],
		Trigger: h.Config.Trigger,
	}

	h.DB.Find(&head, model.Head{Channel: vars["name"]})

	if len(head) == 0 {
		content.Error = "Could not find any quotes"
	} else {
		for _, hx := range head {
			var lines []model.Line
			h.DB.Model(&hx).Related(&lines)

			if hx.Title == "" {
				hx.Title = fmt.Sprintf(`%s "%s"`,
					lines[0].Author, lines[0].Content)
			}

			out = append(out, hx)
		}

		content.Instances = out
	}

	err := templates.ExecuteTemplate(w, "channelindex", content)

	if err != nil {
		log.Println(err)

		fmt.Fprintf(w, "An internal server error occurred: %s", err)
	}
}
