package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

var port = flag.String("port", "8080", "http server port")

func main() {
	flag.Parse()
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Home page\n")
}
