package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"simxd.com/books/config/db"
	myrooter "simxd.com/books/router"
)

func main(){
	db.InitDB()
	router := myrooter.NewRooter(mux.NewRouter())
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300, // Maximum value not ignored by any of major browsers
		AllowCredentials: true,
		Debug:            true,

	}).Handler(router)
	log.Fatal(http.ListenAndServe("localhost:8080", handler))
}