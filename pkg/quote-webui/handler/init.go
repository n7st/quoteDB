package handler

import (
	"html/template"

	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	"github.com/n7st/quoteDB/util"
)

type Handler struct {
	DB     *gorm.DB
	Config *util.Config
}

var templates = template.Must(template.ParseGlob("view/*"))

func NewHandler(db *gorm.DB, c *util.Config) *Handler {
	return &Handler{DB: db, Config: c}
}

func (h *Handler) Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/view/{id}", h.ViewHandler)
	r.HandleFunc("/channel/{name}", h.IndexHandler)

	return r
}
