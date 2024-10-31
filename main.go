package main

import (
	"fortress/handler"
	"fortress/repo"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	_, err := repo.GetConnection()
	if err != nil {
		log.Fatalln(err)
	}

	r := mux.NewRouter()
	r.Use(jsonContentTypeMiddleware)

	r.HandleFunc("/", handler.RootHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/users", handler.CreateUser).Methods("POST")

	log.Println("Listening on port 8085")
	err = http.ListenAndServe(":8085", r)
	if err != nil {
		panic(err)
	}
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
