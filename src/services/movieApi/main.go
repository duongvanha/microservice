package main

import (
	"log"
	micro_app "microservice/src/gopkg/core/microApp"
	"microservice/src/gopkg/core/transport/transhttp"
	micro_models "microservice/src/gopkg/models"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	transhttp.RespondJSON(w, http.StatusOK, &micro_models.Movie{
		Id:              1,
		ContactId:       "2",
		Tags:            "3",
		Status:          4,
		Source:          5,
		TicketCreatedAt: 6,
		TicketUpdatedAt: 7,
		CreatedAt:       8,
		UpdatedAt:       9,
		AppCode:         "10",
		Platform:        "11",
	})
}

func main() {

	service := micro_app.NewHTTPApp()

	service.HandleFunc("/", helloWorldHandler)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
