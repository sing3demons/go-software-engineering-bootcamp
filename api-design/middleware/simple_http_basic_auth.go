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

type Logger struct {
	Handler http.Handler
}

func logMiddleware(Handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		Handler.ServeHTTP(w, r)
		log.Printf("Server http middleware: %s %s %s %s", r.RemoteAddr, r.Method, r.URL, time.Since(start))
	}
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.Handler.ServeHTTP(w, r)
	log.Printf("Server http middleware: %s %s %s %s", r.RemoteAddr, r.Method, r.URL, time.Since(start))
}

func main() {
	mux := http.NewServeMux()
	logMux := Logger{Handler: mux}
	srv := http.Server{
		Addr:    ":2565",
		Handler: logMux,
	}

	mux.HandleFunc("/users", AuthMiddleware(usersHandle))
	mux.HandleFunc("/health", healthHandler)

	log.Println("Server started as :2565")
	log.Fatal(srv.ListenAndServe())
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		log.Println("auth:", u, p, ok)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`can't parse the basic auth`))
			return
		}

		if u != "api" || p != "api" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`Username or Password incorrect.`))
			return
		}

		next.ServeHTTP(w, r)
	}
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

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
