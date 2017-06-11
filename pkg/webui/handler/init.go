package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/view/{id}", h.ViewHandler)
	r.HandleFunc("/channel/{name}", h.IndexHandler)

	return r
}
