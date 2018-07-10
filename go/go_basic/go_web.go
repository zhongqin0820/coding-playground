package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	for k, v := range r.Form {
		fmt.Println("key:", k, " val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello...")
}

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe("localhost:9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
