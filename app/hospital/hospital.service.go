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
}

func ProvideHospitalService() HospitalService {
	return HospitalService{
		BaseURL:       os.Getenv("JSC_URL"),
		Authorization: os.Getenv("JSC_TOKEN"),
		ContentType:   "application/json; charset=UTF-8",
	}
}

func (s *HospitalService) GetHospital() (Response, int) {
	url := s.BaseURL + "/rumahsakitumum"
	request, err := http.Get(url)
	if err != nil {
		return Response{}, 500
	}

	data, _ := ioutil.ReadAll(request.Body)

	response := Response{}
	json.Unmarshal(data, &response)

	return response, request.StatusCode
}
