package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var DB *sql.DB

func main() {
	r := mux.NewRouter()
	rout(r)
	DB = conn()
	createtable(DB)
	fmt.Println("running in port:8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
