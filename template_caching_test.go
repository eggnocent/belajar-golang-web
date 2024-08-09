package belajargolangweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "./templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "hello template caching")
}

func TestTemplaceCaching(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
