// The handler package contains controllers for displaying web pages.
package handler

import (
	"html/template"

	"github.com/n7st/quoteDB/util"

	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
)

// Handler{} provides the base for a HTTP controller.
type Handler struct {
	DB     *gorm.DB
	Config *util.Config
}

// Parse all the templates ready for display.
var templates = template.Must(template.ParseGlob("view/*"))

// NewHandler() Creates a new Handler{}.
func NewHandler(db *gorm.DB, c *util.Config) *Handler {
	return &Handler{DB: db, Config: c}
}

// Router() sets up HTTP routes.
func (h *Handler) Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/view/{id}", h.ViewHandler)
	r.HandleFunc("/channel/{name}", h.ChannelIndexHandler)

	return r
}
