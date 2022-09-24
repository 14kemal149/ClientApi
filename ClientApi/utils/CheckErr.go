package utils

/*
Ahmet Kemal Bayraktar
23.09.2022
ClientApi uygulamasi icin basit bir hata yonetimi

*/

import (
	"log"
)

func CheckErr(err error) { //Basit bir hata yonetimi
	if err != nil {
		log.Fatal(err)
	}
}
