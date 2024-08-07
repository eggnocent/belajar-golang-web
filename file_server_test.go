package belajargolangweb

import (
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileserver := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", fileserver)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
