package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		profile := Profile{"Alex", []string{"snowboarding", "programming"}}

		js, err := json.Marshal(profile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	http.Handle("/", r)

	port := "8081"

	log.Println("Starting HTTP service at " + port)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}

}
