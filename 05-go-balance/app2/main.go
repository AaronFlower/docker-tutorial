package main

import (
	"log"
	"net/http"
)

func helloFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello 2"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloFunc)
	log.Fatal(http.ListenAndServe(":80", mux))
}
