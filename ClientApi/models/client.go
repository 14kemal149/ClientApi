package models

type client struct {
	ID        int    `json : "ID"`
	firstname string `json : "Firstname"`
	lastname  string `json : "Lastname"`
}
