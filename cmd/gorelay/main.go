package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"gorelay/internal/gorelay"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println("default: ", r.Form)
	fmt.Println("path ", r.URL.Path)
	fmt.Println("param: ", r.Form["test_param"])

	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Golan Webserver Working!")
}


func main() {
	serv := server.NewServer()

	go serv.Serve()
	defer serv.Close()

	http.Handle("/socket.io/", serv)
	log.Println("Serving at")
	log.Fatal(http.ListenAndServe(":8000", nil))
}