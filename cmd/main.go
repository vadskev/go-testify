package main

import (
	"fmt"
	"net/http"

	"github.com/vadskev/go-testify/internal/handler"
)

func main() {
	address := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/cafe", handler.MainHandle)

	srv := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	fmt.Printf("Starting server on %s\n", address)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Printf("failed to listen and serve: %s\n", err)
	}
}
