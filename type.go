package main

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// type DBconnection struct {
// 	PG *sql.DB
// }
