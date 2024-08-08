package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TemplateData(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template dari map",
		"Nama":  "Egi Wiratama",
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateData(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateAction(writer http.ResponseWriter, request *http.Request) {
	t := template.New("name.gohtml").Funcs(template.FuncMap{
		"upper": strings.ToUpper,
	})
	t = template.Must(t.ParseFiles("./templates/name.gohtml"))

	data := map[string]interface{}{
		"Title":    "Template dari map",
		"Nama":     "Crarrson",
		"IsMember": true,
		"Items":    []string{"item 1", "item 2", "item 3"},
	}
	t.ExecuteTemplate(writer, "name.gohtml", data)
}
