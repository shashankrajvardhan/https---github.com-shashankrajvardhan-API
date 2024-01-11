package main

import "github.com/gorilla/mux"

func rout(r *mux.Router) {
	r.HandleFunc("/home/fetch", getUsers).Methods("GET")
	// r.HandleFunc("/home/crete", getUsers).Methods("POST")
	// r.HandleFunc("/home/update", getUsers).Methods("PUT")
	// r.HandleFunc("/home/delete", getUsers).Methods("DELETE")
}
