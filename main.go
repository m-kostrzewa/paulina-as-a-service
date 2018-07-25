package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("Starting serving on localhost:8080")
	r := mux.NewRouter()
	r.HandleFunc("/say", Say)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.ListenAndServe(":8080", r)
}

func Say(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler")

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	text := r.PostFormValue("text")
	fmt.Println("would say", string(text))
}
