package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello kubernetes %q", html.EscapeString(r.URL.Path))
	})
	
	log.Fatal(http.ListenAndServe(":9200", nil))
}
