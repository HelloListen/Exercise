package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// func responseHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		fmt.Println(r.URL.Path)
// 	}
// }

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %q", r.Method, html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8899", nil))
}
