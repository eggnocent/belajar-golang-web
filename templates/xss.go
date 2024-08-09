package templates

import (
	"html/template"
	"net/http"
)

var myTemplates = template.Must(template.ParseFiles("post.gohtml"))

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Auto escape",
		"Body":  "<h3>in body</h3>",
	})
}
