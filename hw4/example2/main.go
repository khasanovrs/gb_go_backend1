package main

import (
	"fmt"
	"net/http"
	"time"
)

type helloHandler struct {
	subject string
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, %s!", h.subject)
	if err != nil {
		return
	}
}

func main() {
	worldHandler := &helloHandler{"World"}
	roomHandler := &helloHandler{"Mark"}

	http.Handle("/world", worldHandler)
	http.Handle("/room", roomHandler)

	srv := &http.Server{
		Addr:         ":80",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return
	}
}
