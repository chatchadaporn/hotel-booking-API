package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"booking-service/handler/booking"
	"booking-service/handler/search"
)

func initData() {

}

func main() {
	search.RoomDetail = "hello"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/search/{id}", search.Handler).Methods("GET")
	router.HandleFunc("/searchs", search.Handler).Methods("GET")
	router.HandleFunc("/booking", booking.Handler).Methods("Post")
	log.Fatal(http.ListenAndServe(":1121", router))

}
