package belajargolangweb

import (
	"net/http"
)

//var templates embed.FS

//var myTemplates = template.Must(template.ParseFS(templates, "./templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "hello template caching")
}
