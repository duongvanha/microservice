package bilutv

import "testing"

func TestGateway_GetMovie(t *testing.T) {
	response, err := Gateway{}.GetMovie(struct {
		Id int64
		Ep int64
	}{Id: 13007, Ep: 167247})

	println(response, err)
}
