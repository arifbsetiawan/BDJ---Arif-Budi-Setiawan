package subdistrict

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type SubdistrictService struct {
	BaseURL       string
	Authorization string
	ContentType   string
}

func ProvideSubdistrictService() SubdistrictService {
	return SubdistrictService{
		BaseURL:       os.Getenv("JSC_URL"),
		Authorization: os.Getenv("JSC_TOKEN"),
		ContentType:   "application/json; charset=UTF-8",
	}
}

func (s *SubdistrictService) GetSubdistrict() (Response, int) {
	url := s.BaseURL + "/kelurahan"
	request, err := http.Get(url)
	if err != nil {
		return Response{}, 500
	}

	data, _ := ioutil.ReadAll(request.Body)

	response := Response{}
	json.Unmarshal(data, &response)

	return response, request.StatusCode
}
