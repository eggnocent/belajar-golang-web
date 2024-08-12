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

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) { // Ubah request menjadi *http.Request
	defer func() {
		if err := recover(); err != nil { // Pastikan untuk menangkap panic
			fmt.Println("ERROR:", err)
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "error: %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request) // Tidak perlu menggunakan &request
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo Executed")
		fmt.Fprint(writer, "Hello Foo")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Panic Executed")
		panic("ups panic")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}
	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	// Jalankan server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
