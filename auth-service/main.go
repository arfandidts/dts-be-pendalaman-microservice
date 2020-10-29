package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arfandidts/dts-be-pengenalan-microservice/auth-service/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/admin-auth", http.HandlerFunc(handler.ValidateAuth))

	fmt.Printf("Auth service listen on :5001")
	log.Panic(http.ListenAndServe(":5001", router))
}
