package main

import (
	"flag"
	"fmt"
	//	"io"
	"log"
	"net/http"

	"handlers"
)

var port = flag.String("port", "8080", "http server port")

func main() {
	flag.Parse()
	//	http.HandleFunc("/", indexHandler)
	handler := &handlers.Handler{}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), handler))
}

//func indexHandler(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, "Home page\n")
//}
