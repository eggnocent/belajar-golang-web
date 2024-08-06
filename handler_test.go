package belajargolangweb

import (
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(writter http.ResponseWritter, request *http.Request) {
		// logic
		// writter untuk mengembalikan response

	}
}
