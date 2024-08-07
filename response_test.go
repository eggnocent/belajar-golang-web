package belajargolangweb

import (
	"fmt"
	"net/http"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	// parsing otomatis
	request.PostFormValue("firstname")

	firstName := request.PostForm.Get("firstName")
	lastName := request.PostForm.Get("lastName")
	fmt.Fprintf(writer, "%s %s", firstName, lastName)
}
