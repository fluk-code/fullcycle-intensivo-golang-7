package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"github.com/mattn/go-sqlite3"
)

type User struct {
	ID int
	Name string 
	Email string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", listUserHandler)
	http.ListenAndServe(":3333", mux)
}

func listUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("List of users")))

	db, err := sql.Open("sqlite3", "users.db")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}