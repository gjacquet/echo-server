package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func prettyPrintJSON(reader io.Reader) (string, error) {
	decoder := json.NewDecoder(reader)

	var body interface{}
	if err := decoder.Decode(&body); err != nil {
		return "", err
	}

	bytes, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func extractBody(reader io.Reader) (string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func echo(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v", r.Method, r.URL, r.Proto)
	log.Printf("Host: %v", r.Host)
	log.Print("Headers:")
	for name, values := range r.Header {
		log.Printf("  %v: %v", name, strings.Join(values, ", "))
	}

	if body := r.Body; body != nil {
		var bodyString string
		var err error

		contentType := r.Header.Get("Content-Type")
		switch contentType {
		case "application/json":
			bodyString, err = prettyPrintJSON(body)
		default:
			bodyString, err = extractBody(body)
		}

		if err != nil {
			log.Printf("Error: %v", err)

			http.Error(w, "Could not read body", http.StatusInternalServerError)
			return
		}

		if err != nil {
			log.Printf("Error: %v", err)

			http.Error(w, "Could not read body", http.StatusInternalServerError)
			return
		}

		if bodyString != "" {
			log.Print("Body:")
			log.Print(bodyString)
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
