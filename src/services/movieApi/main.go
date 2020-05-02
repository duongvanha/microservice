package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
	micro_app "microservice/src/gopkg/core/microApp"
	"microservice/src/gopkg/core/transport/transhttp"
	micro_models "microservice/src/gopkg/models"
	MovieRepository "microservice/src/gopkg/services"
	"net/http"
	"time"
)

type handler struct {
	Client MovieRepository.MovieRepositoryService
}

func (h handler) helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	res, err := h.Client.Create(context.Background(), &micro_models.Movie{
		Id:              time.Now().Unix(),
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

	if err != nil {
		log.Error("error", err)
	} else {
		transhttp.RespondJSON(w, http.StatusOK, res)
	}

}

func main() {

	service := micro_app.NewWebApp()

	consignmentService := micro_app.NewService(micro.Name("go.haduong.service.build_movie"))

	handler := handler{
		Client: MovieRepository.NewMovieRepositoryService("go.haduong.service.build_movie", consignmentService.Client()),
	}

	service.HandleFunc("/", handler.helloWorldHandler)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
