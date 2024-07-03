package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})
	createHandlerFromYaml(mux)
	createHandlerFromJson(mux)
	http.ListenAndServe(":9090", mux)
}
