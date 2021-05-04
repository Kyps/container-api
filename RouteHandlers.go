package main

import (
	"net/http"
)

// Route for fetching all files in database. Only "GET" method is accepted, other methods return 404 - Not Found
func handleGetAll(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		result, err := conf.GetAllFiles()
		if err != nil {
			ErrorInternal(w, req, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	default:
		ErrorNotFound(w, req, nil)
	}
}

// Route for creating and deleting signatures from containers. Only "PUT" and "DELETE" methods are accepted,
// other methods return 404 - Not Found
func handleSignatures(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "PUT":
		conf.AddSignatureController(w, req)
	case "DELETE":
		conf.DeleteSignatureController(w, req)
	default:
		ErrorNotFound(w, req, nil)
	}
}

// Route index endpoint for creating new container. Only "POST" method is accepted, other methods return 404 - Not Found
func handleIndex(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		conf.PostController(w, req)
	default:
		ErrorNotFound(w, req, nil)
	}

}
