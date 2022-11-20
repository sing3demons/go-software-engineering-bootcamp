package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	users := []User{{ID: 1, Name: "sing", Age: 21}}
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
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
				w.Write([]byte(err.Error()))
			}

			user := User{}

			err = json.Unmarshal(body, &user)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}

			w.Write(body)
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	log.Println("Server started as :2565")
	log.Fatal(http.ListenAndServe(":2565", nil))
}
