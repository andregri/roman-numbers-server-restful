package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/andregri/roman-numbers-server-restful/romanNumerals"
)

func main() {
	// handle main route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")

		request := urlPathElements[1]
		input := urlPathElements[2]

		// If it is a GET request and the syntax is correct
		if request == "roman_number" {
			number, _ := strconv.Atoi(strings.TrimSpace(input))

			// If the resource is not in the list, send a Not Found status
			if number == 0 || number > 10 {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not found"))
			} else {
				fmt.Fprintf(w, "%q", html.EscapeString(romanNumerals.Numerals[number]))
			}
		} else {
			// For all other requests, tell the client sent a Bad Request
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
		}
	})

	// Create a server and run it on port 8000
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
