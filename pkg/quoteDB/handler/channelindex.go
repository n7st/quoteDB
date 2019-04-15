// The handler package contains controllers for displaying web pages.
package handler

import (
	"fmt"
	"log"
	"net/http"

	"git.netsplit.uk/mike/quoteDB/pkg/quoteDB/model"

	"github.com/gorilla/mux"
)

// ChannelIndexContent{} contains template variables.
type ChannelIndexContent struct {
	Instances []model.Head
	Channel   string
	Error     string
	Trigger   string
}

// ChannelIndexHandler() displays a page with a list of quotes for the given
// {channel}.
func (h *Handler) ChannelIndexHandler(w http.ResponseWriter, r *http.Request) {
	var (
		channel model.Channel
		heads   []model.Head
		out     []model.Head
	)

	vars := mux.Vars(r)
	content := &ChannelIndexContent{
		Channel: vars["name"],
		Trigger: h.Config.Trigger,
	}

	h.DB.Find(&channel, model.Channel{Name: vars["name"]})
	h.DB.Model(&channel).Related(&heads)

	for _, hx := range heads {
		var lines []model.Line
		h.DB.Model(&hx).Related(&lines)

		if hx.Title == "" {
			hx.Title = fmt.Sprintf(`%s "%s"`,
				lines[0].Author, lines[0].Content)
		}

		out = append(out, hx)
	}

	content.Instances = out

	err := templates.ExecuteTemplate(w, "channelindex", content)

	if err != nil {
		log.Println(err)

		fmt.Fprintf(w, "An internal server error occurred: %s", err)
	}
}
