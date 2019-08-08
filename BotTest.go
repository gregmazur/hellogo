package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":8080", router))
}


func test(w http.ResponseWriter, r *http.Request){
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	log.Println("------recieved " + r.Method + " req-----")
	log.Println(r.Header)
	log.Println(string(reqBody))
	log.Println("------END-----")
}


