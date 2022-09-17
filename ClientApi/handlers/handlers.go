package handlers

import (
	"net/http"
)

// http get client
func getclienthandler(w http.ResponseWriter, r *http.Request) {

	clients := make(map[string]string)

	//clients := loaduser()

	for _, client := range clients {

	}

}

// http post client
// http get allclient
// http delete client
// http put client
