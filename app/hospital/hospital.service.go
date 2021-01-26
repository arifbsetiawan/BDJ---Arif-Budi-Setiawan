package hospital

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type HospitalService struct {
	BaseURL       string
	Authorization string
	ContentType   string
	Client        *http.Client
}

func ProvideHospitalService() HospitalService {
	return HospitalService{
		BaseURL:       os.Getenv("JSC_URL"),
		Authorization: os.Getenv("JSC_TOKEN"),
		ContentType:   "application/json; charset=UTF-8",
		Client:        &http.Client{},
	}
}

func (s *HospitalService) GetHospital() (Response, int) {
	url := s.BaseURL + "/rumahsakitumum"
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
