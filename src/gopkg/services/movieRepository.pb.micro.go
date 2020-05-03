// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: services/movieRepository.proto

package micro_services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	models "microservice/src/gopkg/models"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for MovieRepository service

type MovieRepositoryService interface {
	Create(ctx context.Context, in *models.Movie, opts ...client.CallOption) (*models.Movie, error)
}

type movieRepositoryService struct {
	c    client.Client
	name string
}

func NewMovieRepositoryService(name string, c client.Client) MovieRepositoryService {
	return &movieRepositoryService{
		c:    c,
		name: name,
	}
}

func (c *movieRepositoryService) Create(ctx context.Context, in *models.Movie, opts ...client.CallOption) (*models.Movie, error) {
	req := c.c.NewRequest(c.name, "MovieRepository.Create", in)
	out := new(models.Movie)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MovieRepository service

type MovieRepositoryHandler interface {
	Create(context.Context, *models.Movie, *models.Movie) error
}

func RegisterMovieRepositoryHandler(s server.Server, hdlr MovieRepositoryHandler, opts ...server.HandlerOption) error {
	type movieRepository interface {
		Create(ctx context.Context, in *models.Movie, out *models.Movie) error
	}
	type MovieRepository struct {
		movieRepository
	}
	h := &movieRepositoryHandler{hdlr}
	return s.Handle(s.NewHandler(&MovieRepository{h}, opts...))
}

type movieRepositoryHandler struct {
	MovieRepositoryHandler
}

func (h *movieRepositoryHandler) Create(ctx context.Context, in *models.Movie, out *models.Movie) error {
	return h.MovieRepositoryHandler.Create(ctx, in, out)
}