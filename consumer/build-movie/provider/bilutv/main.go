package bilutv

import (
	"encoding/json"
	"fmt"
	"github.com/duongvanha/microservice/consumer/build-movie/provider"
	"github.com/imroc/req"
	"regexp"
)

var re = regexp.MustCompile("sources:\\[(.*)],")

type Gateway struct{}

func (Gateway) GetMovie(dto provider.MovileDto) (movie *provider.Movie, err error) {
	r := req.New()
	response, err := r.Post("https://bilutv.org/ajax/player", req.Param{
		"id": dto.Id,
		"ep": dto.Ep,
	})

	if err != nil {
		return
	}

	text, err := response.ToString()

	if err != nil {
		return
	}

	match := re.FindStringSubmatch(text)

	if len(match) != 2 {
		return nil, provider.ParseError
	}

	var sources []source

	content := fmt.Sprintf(`[%v]`, match[1])

	err = json.Unmarshal([]byte(content), &sources)

	if err != nil {
		return
	}

	var chaps []provider.Chap

	for _, source := range sources {
		chaps = append(chaps, provider.Chap{
			Url:      source.File,
			Type:     source.Type,
			Quantity: source.Label,
		})
	}

	return &provider.Movie{
		Chaps: chaps,
	}, nil
}

type source struct {
	File  string `json:"file"`
	Label string `json:"label"`
	Type  string `json:"type"`
}
