package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang Sample, Selamat Datang!\nApp Version: 8\nBelajar Jenkins Pipeline Automation\nProject GOAPP\n\n<h1>MANTAPPPPP</h1>")
}

func envPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of environment variable:")
	for _, env := range os.Environ() {
		fmt.Fprintf(w, "- "+env+"\n")
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/env", envPage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequests()
}