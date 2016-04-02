package main

import (
	"fmt"
	"net/http"
	"html"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.Index(html.EscapeString(r.URL.Path), "/g/") == 0 {
			name := strings.Replace(strings.Replace(strings.Split(html.EscapeString(r.URL.Path), "g/")[1], "\"", "", -1), "/", " ", -1)
			fmt.Fprintf(w, "Goodbye %s!", name)
		} else {
			name := strings.Replace(strings.Replace(html.EscapeString(r.URL.Path), "\"", "", -1), "/", " ", -1)
			fmt.Fprintf(w, "Hello%s!", name)
		}
	})
	http.ListenAndServe(":8080", nil)
}
