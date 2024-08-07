package belajargolangweb

import (
	"fmt"
	"net/http"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	nama := request.URL.Query().Get("name")
	if nama == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Println(writer, "name is empty")
	} else {
		fmt.Fprintf(writer, "hello %s", nama)
	}
}
