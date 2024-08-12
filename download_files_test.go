package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	file := request.URL.Query().Get("file")

	if file == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Bad Request")
		return
	}

	// Perbaiki kesalahan penulisan "Content-Disposition"
	writer.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")

	http.ServeFile(writer, request, "./resources/"+file)
}

func TestDownloadFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/download", DownloadFile)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
