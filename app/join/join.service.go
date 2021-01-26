package join

import (
	"errors"
	"jsc-join-api/app/hospital"
	"jsc-join-api/app/subdistrict"
	"net/http"
)

type JoinService struct {
	subdistrictService subdistrict.SubdistrictService
	hospitalService    hospital.HospitalService
}

func ProvideJoinService(
	SubdistrictService subdistrict.SubdistrictService,
	HospitalService hospital.HospitalService,
) JoinService {
	return JoinService{
		subdistrictService: SubdistrictService,
		hospitalService:    HospitalService,
	}
}

func (s *JoinService) Find() (Response, error) {
	hospitals, code := s.hospitalService.GetHospital()
	if code != http.StatusOK {
		return Response{}, errors.New("Cannot connect to hospital")
	}

	subdistricts, code := s.subdistrictService.GetSubdistrict()
	if code != http.StatusOK {
		return Response{}, errors.New("Cannot connect to subdistrict")
	}

	joins := []Join{}

	for _, data := range hospitals.Data {
		var subData subdistrict.Subdistrict
		for _, sub := range subdistricts.Data {
			if sub.SubdistrictCode == data.SubdistictCode {
				subData = sub
			}
		}
		join := Join{
			ID:           data.ID,
			HospitalName: data.HospitalName,
			HospitalType: data.HospitalType,
			Location: Location{
				Latitude:  data.Location.Latitude,
				Longitude: data.Location.Longitude,
			},
			Address:    data.Location.Address,
			PostalCode: data.PostalCode,
			Phone:      data.Phone,
			Fax:        data.Fax,
			Website:    data.Website,
			Email:      data.Email,
			Subdistrict: Subdistrict{
				SubdistictCode: subData.SubdistrictCode,
				SubdistictName: subData.SubdistrictName,
			},
			District: District{
				DistrictCode: subData.DistrictCode,
				DistrictName: subData.DistrictName,
			},
			City: City{
				CityCode: subData.CityCode,
				CityName: subData.CityName,
			},
		}
		joins = append(joins, join)
	}

	response := Response{}
	response.Status = "success"
	response.Count = len(joins)
	response.Data = joins

	return response, nil
}
