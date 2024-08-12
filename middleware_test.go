package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("before execute handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("after execute handler")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()

	// Definisikan handler yang akan dipanggil
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello middleware")
	})

	// Buat instance LogMiddleware
	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: logMiddleware,
	}

	// Jalankan server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	err = server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
