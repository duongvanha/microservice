package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"math/rand"
	micro_app "microservice/src/gopkg/core/microApp"
	"microservice/src/gopkg/core/transport/transhttp"
	micro_models "microservice/src/gopkg/models"
	MovieRepository "microservice/src/gopkg/services"
	"net/http"
	"sync"
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
		log.Infof("error", err)
	} else {
		transhttp.RespondJSON(w, http.StatusOK, res)
	}

}

func RoundRobin() func([]*registry.Service) selector.Next {
	i := rand.Int()
	return func(services []*registry.Service) selector.Next {
		nodes := make([]*registry.Node, 0, len(services))

		for _, service := range services {
			nodes = append(nodes, service.Nodes...)
		}

		println("here", len(nodes))

		var mtx sync.Mutex

		return func() (*registry.Node, error) {
			if len(nodes) == 0 {
				return nil, selector.ErrNoneAvailable
			}

			mtx.Lock()
			node := nodes[i%len(nodes)]
			i++
			mtx.Unlock()

			return node, nil
		}
	}
}

func main() {

	service := micro_app.NewHTTPApp()

	consignmentService := micro.NewService(
		micro.Name("shippy.service.consignment"),
		micro.Version("latest"),
		micro.Selector(selector.NewSelector(
			selector.SetStrategy(RoundRobin()),
		)),
	)

	handler := handler{
		Client: MovieRepository.NewMovieRepositoryService("shippy.service.consignment", consignmentService.Client()),
	}

	service.HandleFunc("/", handler.helloWorldHandler)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
