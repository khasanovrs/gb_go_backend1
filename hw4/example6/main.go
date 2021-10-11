package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
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

// UploadHandler /**
type UploadHandler struct {
	HostAddr  string
	UploadDir string
}

func (h *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			http.Error(w, "Unable to close file", http.StatusInternalServerError)
			return
		}
	}(file)

	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	separator := string(os.PathSeparator)
	filePath := h.UploadDir + separator + header.Filename

	err = ioutil.WriteFile(filePath, data, 0o600)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "File %s has been successfully uploaded\n", header.Filename)
	if err != nil {
		return
	}

	fileLink := h.HostAddr + "/" + header.Filename
	_, err = fmt.Fprintln(w, fileLink)
	if err != nil {
		return
	}
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
	uploadHandler := &UploadHandler{
		HostAddr:  "localhost:8080",
		UploadDir: "upload",
	}
	http.Handle("/upload", uploadHandler)
	http.Handle("/", handler)

	srv := &http.Server{
		Addr:         ":80",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			return
		}
	}()

	dirToServe := http.Dir(uploadHandler.UploadDir)

	fs := &http.Server{
		Addr:         ":8080",
		Handler:      http.FileServer(dirToServe),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := fs.ListenAndServe()
	if err != nil {
		return
	}
}
