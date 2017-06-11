package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"fmt"
	"github.com/n7st/quoteDB/model"
)

type IndexContent struct {
	Instances []model.Head
	Channel   string
	Error     string
}

type headWithLines struct {
	Head  model.Head
	Lines []model.Line
}

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	var (
		head []model.Head
		out  []model.Head
	)

	vars := mux.Vars(r)
	content := &IndexContent{Channel: vars["name"]}

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

	err := templates.ExecuteTemplate(w, "channel_index", content)

	if err != nil {
		log.Println(err)
	}
}
