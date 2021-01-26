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
	Client        *http.Client
}

func ProvideSubdistrictService() SubdistrictService {
	return SubdistrictService{
		BaseURL:       os.Getenv("JSC_URL"),
		Authorization: os.Getenv("JSC_TOKEN"),
		ContentType:   "application/json; charset=UTF-8",
		Client:        &http.Client{},
	}
}

func (s *SubdistrictService) GetSubdistrict() (Response, int) {
	url := s.BaseURL + "/kelurahan"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Response{}, 500
	}

	req.Header.Set("Authorization", s.Authorization)
	request, err := s.Client.Do(req)
	if err != nil {
		return Response{}, 500
	}
	data, _ := ioutil.ReadAll(request.Body)

	response := Response{}
	json.Unmarshal(data, &response)

	return response, request.StatusCode
}
