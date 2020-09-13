package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, h *http.Request) {
	w.Write([]byte("hello kir"))

}

func showSnippet(w http.ResponseWriter, h *http.Request) {
	w.Write([]byte("Display a specific snippet"))

}

func createSnippet(w http.ResponseWriter, h *http.Request) {
	w.Write([]byte("create a snippet"))

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homepage)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fmt.Println("Server is running")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
