package belajargolangweb

import (
	"html/template"
	"net/http"
	"testing"
)

var myTemplatesUpload = template.Must(template.ParseFiles("templates/upload_form.html"))

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	err := myTemplatesUpload.ExecuteTemplate(writer, "upload_form.html", nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/form", UploadForm)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
