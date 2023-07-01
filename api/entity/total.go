package entity

type GetTotalResult struct {
	Provinces    int `json:"provinces"`
	Cities       int `json:"cities"`
	Subdistricts int `json:"subdistricts"`
	Villages     int `json:"villages"`
}
