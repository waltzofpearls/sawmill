package api

import (
	"log"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		log.Fatalf("Internal error: %s", err.Error())
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}
