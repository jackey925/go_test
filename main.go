package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello kubernetes %v\n", r.Header)
	})
	log.Fatal(http.ListenAndServe(":9200", nil))
}


