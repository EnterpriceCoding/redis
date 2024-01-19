package main

import (
	"fmt"
	"net/http"
)

func handleGuestPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"hello server!")
}