package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ERROR: %v", err)
		return
	}
	fmt.Fprintf(w, "Form Submit Successful\n")
	fName := r.FormValue("fName")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %v\n", fName)
	fmt.Fprintf(w, "Address: %v\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.!!!", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello, world!\n")
	fmt.Fprintf(w, "2 - Hello, world!")
}

func main() {
	PORT := 8080
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port %v\n", PORT)

	err := http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
	if err != nil {
		log.Fatal(err)
	}
}
