package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love Go!\n\n")

	if db == nil {
		http.Error(w, "database not initialized", http.StatusInternalServerError)
		return
	}

	var message string
	err := db.QueryRow("SELECT message FROM greetings LIMIT 1").Scan(&message)
	if err != nil {
		http.Error(w, fmt.Sprintf("query failed: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", message)
}

func main() {
	godotenv.Load()
	InitDatabase()

	defer CloseDB()

	http.HandleFunc("/", handler)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
