package subdistrict

type Response struct {
	Status string        `json:"status"`
	Count  int           `json:"count"`
	Data   []Subdistrict `json:"data"`
}

type Subdistrict struct {
	ProvinceCode    int    `json:"kode_provinsi"`
	ProvinceName    string `json:"nama_provinsi"`
	CityCode        int    `json:"kode_kota"`
	CityName        string `json:"nama_kota"`
	DistrictCode    int    `json:"kode_kecamatan"`
	DistrictName    string `json:"nama_kecamatan"`
	SubdistrictCode int    `json:"kode_kelurahan"`
	SubdistrictName string `json:"nama_kelurahan"`
}
