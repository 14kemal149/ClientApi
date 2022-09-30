package models

/*
Ahmet Kemal Bayraktar
23.09.2022
ClientApi uygulamasi icin musteri modelinin belirlenmesi

*/

type Client struct {
	ID        int    `json : "ID"`
	Firstname string `json : "Firstname"`
	Lastname  string `json : "Lastname"`
}
