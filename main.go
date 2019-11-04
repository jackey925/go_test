package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello kubernetes %q\n", html.EscapeString(r.URL.Host))
	})
	log.Fatal(http.ListenAndServe(":9200", nil))
}


