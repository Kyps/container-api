package main

import (
	"log"
	"net/http"
)

// Respord with a "404 - Not Found"
func ErrorNotFound(w http.ResponseWriter, req *http.Request, err error) {
	log.Println(err)
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

// Respond with a "500 - Internal Error"
func ErrorInternal(w http.ResponseWriter, req *http.Request, err error) {
	log.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
