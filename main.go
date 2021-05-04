package main

import (
	"log"
	"net/http"
	"os"
)

var conf Config

// Before server starts, get directories from configuration file
// create necessary directories if they do not exist
func init() {
	conf = GetConfig("conf.json")
	if _, err := os.Stat(conf.DatabaseDir); os.IsNotExist(err) {
		os.Mkdir(conf.DatabaseDir, 0755)
	}
	if _, err := os.Stat(conf.TempDir); os.IsNotExist(err) {
		os.Mkdir(conf.TempDir, 0755)
	}
}

// Starting HTTP server and declare endpoints
func main() {
	mux := http.NewServeMux()

	indexHandler := http.HandlerFunc(handleIndex)
	getAllHandler := http.HandlerFunc(handleGetAll)
	signaturesHandler := http.HandlerFunc(handleSignatures)
	mux.Handle("/", indexHandler)
	mux.Handle("/getall", getAllHandler)
	mux.Handle("/sign/", signaturesHandler)
	log.Println("Listening...")
	http.ListenAndServe(":"+conf.Port, mux)
}
