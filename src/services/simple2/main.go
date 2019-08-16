package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello word"))
	})

	http.Handle("/", r)

	port := "8080"

	log.Println("Starting HTTP service at " + port)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}

}
