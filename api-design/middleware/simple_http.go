package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User = []User{{ID: 1, Name: "sing", Age: 21}}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Server http middleware: %s %s %s %s", r.RemoteAddr, r.Method, r.URL, time.Since(start))
	}
}

func main() {
	http.HandleFunc("/users", logMiddleware(usersHandle))
	http.HandleFunc("/health", logMiddleware(healthHandler))

	log.Println("Server started as :2565")
	log.Fatal(http.ListenAndServe(":2565", nil))
}

func usersHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		b, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.Write(b)
		return
	}

	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error : %v", err)
			return
		}

		id := len(users) + 1

		user := User{}
		user.ID = id

		err = json.Unmarshal(body, &user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "error : %v", err)
			return
		}

		users = append(users, user)

		// w.Write(body)
		fmt.Fprintf(w, "create user")
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
