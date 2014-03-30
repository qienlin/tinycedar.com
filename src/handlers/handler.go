package handlers

import (
	"log"
	"net/http"
)

type Handler struct {
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("ServeHttp...\n")
}
