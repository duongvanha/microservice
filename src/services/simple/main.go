package main

import (
	"github.com/micro/go-micro/v2/logger"
	micro_app "microservice/src/gopkg/core/microApp"
	"microservice/src/gopkg/core/transport/transhttp"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {

	service := micro_app.NewWebApp()

	service.HandleFunc("/helo", func(w http.ResponseWriter, request *http.Request) {
		transhttp.RespondJSON(w, http.StatusOK, Profile{"Alex", []string{"snowboarding", "programming"}})
	})

	service.HandleFunc("/hello", func(w http.ResponseWriter, request *http.Request) {
		transhttp.RespondJSON(w, http.StatusOK, Profile{"Alex", []string{"snowboarding", "programming"}})
	})

	logger.Infof("service running")

	if err := service.Init(); err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}

}
