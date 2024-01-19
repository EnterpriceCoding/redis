package main

import (
	"log"
	"net/http"
)

func main(){

	http.HandleFunc("/", handleGuestPage)

	if err:=http.ListenAndServe(":8080",nil); err!= nil{
		log.Fatal(err)
	}
}