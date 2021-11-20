package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello okteto, Host:%s", r.Host)
		fmt.Fprintf(os.Stdout, "Path: %s, Host: %s\n", r.URL, r.Host)

	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
