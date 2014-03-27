package main

import (
	"flag"
	"io"
	"log"
	"net/http"
)

var port = flag.String("port", "1718", "http server port")

func main() {
	flag.Parse()
	log.Printf("Http server started successfully on port %s\n", *port)

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Home page\n")
}
