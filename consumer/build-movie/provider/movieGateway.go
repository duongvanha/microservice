package provider

import "errors"

type Gateway interface {
	GetMovie(MovileDto) (Movie, error)
}

type MovileDto struct {
	Id int64
	Ep int64
}

type Movie struct {
	Chaps []Chap
	Name  string
}

type Chap struct {
	Url      string
	Type     string
	Quantity string
}

var ParseError = errors.New("parse error")
