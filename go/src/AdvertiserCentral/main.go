package main

import (
	"AdvertiserCentral/AccountCreation"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main(){
	//AccountCreation.CreateAccount()


	//initEvents()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/createaccount",AccountCreation.CreateAccount).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080",router))

}

func homeLink(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello Advertiser!")
}


