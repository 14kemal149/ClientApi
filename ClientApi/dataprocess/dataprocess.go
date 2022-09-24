package dataprocess

/*
Ahmet Kemal Bayraktar
23.09.2022
ClientApi uygulamasi icin data islemleri

*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/14kemal149/ClientApi/models"
	. "github.com/14kemal149/ClientApi/utils"
)

var clients = make(map[string]models.Client)

// loadusers
func Loadusers() map[string]models.Client { //Musterilerin json dosyasindan nesnelere donusum asaması

	bytes, err := ioutil.ReadFile("json/clients.json")
	CheckErr(err)

	err = json.Unmarshal(bytes, &clients)

	return clients
}

// deleteuser
func Deleteuser(id string) { //json dosyasindan musteri silinmesi

	_ = Loadusers()

	var client models.Client

	Id, _ := strconv.Atoi(id)

	for _, client = range clients {
		if client.ID == Id {
			delete(clients, id)
		}
	}
	return
}

// updateuser
func Updateuser(client models.Client, key string) { //json dosyasindaki bir musterinin guncellenmesi
	_ = Loadusers()

	if _, ok := clients[key]; ok {
		delete(clients, key)
		clients[key] = client

	} else {
		fmt.Printf("Değer bulunamadı : %s", key)
	}
	return
}

// saveuser
func Saveuser(client models.Client) { //Yeni kaydolacak musteri icin data islemleri

	_ = Loadusers()

	stid := strconv.Itoa(client.ID - 1)

	clients[stid] = client

	bytes, err := json.Marshal(clients)
	fmt.Println(56)
	//fmt.Println(string(bytes))
	CheckErr(err)

	f, err := os.OpenFile("json/clients.json", os.O_WRONLY, 0600)
	CheckErr(err)

	defer f.Close()

	_, err = f.Write(bytes)
	CheckErr(err)
	return

}
