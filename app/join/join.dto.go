package join

type Response struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
	Data   []Join `json:"data"`
}

type Join struct {
	ID           int         `json:"id"`
	HospitalName string      `json:"nama_rsu"`
	HospitalType int         `json:"jenis_rsu"`
	Location     Location    `json:"location"`
	Address      string      `json:"alamat"`
	PostalCode   int         `json:"kode_pos"`
	Phone        []string    `json:"telepon"`
	Fax          []string    `json:"faximile"`
	Website      string      `json:"website"`
	Email        string      `json:"email"`
	Subdistrict  Subdistrict `json:"kelurahan"`
	District     District    `json:"kecamatan"`
	City         City        `json:"kota"`
}

type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Subdistrict struct {
	SubdistictCode int    `json:"kode"`
	SubdistictName string `json:"nama"`
}

type District struct {
	DistrictCode int    `json:"kode"`
	DistrictName string `json:"nama"`
}

type City struct {
	CityCode int    `json:"kode_kota"`
	CityName string `json:"nama_kota"`
}
