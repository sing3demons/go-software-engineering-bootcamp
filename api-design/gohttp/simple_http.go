package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Write([]byte(`{"name": "sing", "method": "GET"}`))
			return
		}

		if r.Method == http.MethodPost {
			w.Write([]byte(`{"name": "sing", "method": "POST"}`))
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	log.Println("Server started as :2565")
	log.Fatal(http.ListenAndServe(":2565", nil))
	log.Println("Server stop")
}
