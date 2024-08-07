package belajargolangweb

import (
	"embed"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileserver := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// go:embed resources
var resources embed.FS

func TestFileServerGoEmbed(t *testing.T) {
	fileserver := http.FileServer(http.FS(resources))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
