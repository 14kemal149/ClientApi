package models

/*
Ahmet Kemal Bayraktar
23.09.2022
ClientApi uygulamasi icin musteri modelinin belirlenmesi

*/

type Client struct {
	ID        int    `json : "ID"`
	firstname string `json : "Firstname"`
	lastname  string `json : "Lastname"`
}
