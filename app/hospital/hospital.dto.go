package hospital

type Response struct {
	Status string     `json:"status"`
	Count  int        `json:"count"`
	Data   []Hospital `json:"data"`
}

type Hospital struct {
	ID             int      `json:"id"`
	HospitalName   string   `json:"nama_rsu"`
	HospitalType   string   `json:"jenis_rsu"`
	Location       Location `json:"location"`
	PostalCode     int      `json:"kode_pos"`
	Phone          []string `json:"telepon"`
	Fax            []string `json:"faximile"`
	Website        string   `json:"website"`
	Email          string   `json:"email"`
	CityCode       int      `json:"kode_kota"`
	DistrictCode   int      `json:"kode_kecamatan"`
	SubdistictCode int      `json:"kode_kelurahan"`
	Latitude       float64  `json:"latitude"`
	Longitude      float64  `json:"longitude"`
}

type Location struct {
	Address   string  `json:"alamat"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
