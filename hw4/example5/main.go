package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

// Handler /**
type Handler struct{}

// Employee /**
type Employee struct {
	Name   string  `json:"name" xml:"name"`
	Age    int     `json:"age" xml:"age"`
	Salary float32 `json:"salary" xml:"salary"`
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		name := r.FormValue("name")
		_, err := fmt.Fprintf(w, "Parsed query-param with key \"name\": %s", name)
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var employee Employee

		contentType := r.Header.Get("Content-Type")

		switch contentType {
		case "application/json":
			err := json.NewDecoder(r.Body).Decode(&employee)
			if err != nil {
				http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
				return
			}
		case "application/xml":
			err := xml.NewDecoder(r.Body).Decode(&employee)
			if err != nil {
				http.Error(w, "Unable to unmarshal XML", http.StatusBadRequest)
				return
			}
		default:
			http.Error(w, "Unknown content type", http.StatusBadRequest)
			return
		}

		_, err := fmt.Fprintf(w, "Got a new employee!\nName: %s\nAge: %dy.o.\nSalary: %0.2f\n",
			employee.Name,
			employee.Age,
			employee.Salary,
		)
		if err != nil {
			return
		}
	}
}

func main() {
	handler := &Handler{}
	http.Handle("/", handler)

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
