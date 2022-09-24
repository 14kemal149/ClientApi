package main

/*
Ahmet Kemal Bayraktar
23.09.2022
ClientApi uygulamasi icin main islemler

*/

import (
	"net/http"

	. "github.com/14kemal149/ClientApi/handlers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter() //URL uzerinden handler atamaları
	r.HandleFunc("/clients", Getallclientshandler).Methods("GET")
	r.HandleFunc("/client/{id}", Getclienthandler).Methods("GET")
	r.HandleFunc("/client/{id}", Deleteclienthandler).Methods("DELETE")
	r.HandleFunc("/client/{id}", Putclienthandler).Methods("PUT")
	r.HandleFunc("/client", Postclienthandler).Methods("POST")

	server := &http.Server{
		Addr:    ":9090",
		Handler: r,
	}

	server.ListenAndServe()

}
