package handler

import (
	"html/template"

	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
)

type Handler struct {
	DB *gorm.DB
}

var templates = template.Must(template.ParseGlob("view/*"))

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/view/{id}", h.ViewHandler)
	r.HandleFunc("/channel/{name}", h.IndexHandler)

	return r
}
