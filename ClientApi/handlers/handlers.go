package handlers

/*
Ahmet Kemal Bayraktar
23.09.2022
ClientApi uygulamasi icin handler islemleri

*/

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	. "github.com/14kemal149/ClientApi/dataprocess"
	"github.com/14kemal149/ClientApi/models"
	. "github.com/14kemal149/ClientApi/utils"

	"github.com/gorilla/mux"
)

var clients = make(map[string]models.Client)
var id int = 1

// http get client    .../client/{id}
func Getclienthandler(w http.ResponseWriter, r *http.Request) { //id'si verilen musterinin return edilme islemi
	t := mux.Vars(r)
	key := t["id"]

	clients = Loadusers()
	var client models.Client

	for _, client = range clients {
		if strconv.Itoa(client.ID) == key {
			json.NewEncoder(w).Encode(client)
		}

	}
	return
}

// http post client   .../client
func Postclienthandler(w http.ResponseWriter, r *http.Request) { //yeni kaydolacak musteri icin kayit islemi
	var client models.Client

	err := json.NewDecoder(r.Body).Decode(&client)
	fmt.Println(r.Body)

	fmt.Println(22)
	CheckErr(err)

	client.ID = id
	id++

	Saveuser(client)

	data, err := json.Marshal(clients)
	fmt.Println(43)

	CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
	return
}

// http get allclient  .../clients
func Getallclientshandler(w http.ResponseWriter, r *http.Request) { //butun musterilerin return edilmesi

	clients = Loadusers()

	data, err := json.Marshal(clients)

	CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	return
}

// http delete client   .../client/{id}
func Deleteclienthandler(w http.ResponseWriter, r *http.Request) { //id'si verilen musterinin silinme islemi
	t := mux.Vars(r)
	key := t["id"]

	Deleteuser(key)

	w.WriteHeader(http.StatusOK)
	return
}

// http put client  .../client/{id}
func Putclienthandler(w http.ResponseWriter, r *http.Request) { //id'si verilen musterinin guncellenmesi islemi
	t := mux.Vars(r)
	key := t["id"]

	var client models.Client
	err := json.NewDecoder(r.Body).Decode(&client)
	CheckErr(err)

	Updateuser(client, key)

	w.WriteHeader(http.StatusOK)
	return
}
