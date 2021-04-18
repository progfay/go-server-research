package main

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/users/:id"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "/users/:id")
	})

	mux.HandleFunc(pat.Get("/users/new"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "/users/new")
	})

	mux.HandleFunc(pat.Get("/users/files/*"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "/users/files/*")
	})

	http.ListenAndServe("localhost:1323", mux)
}
