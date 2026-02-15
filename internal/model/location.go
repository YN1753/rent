package model

type Location struct {
	Province     string `json:"province" bson:"province"`
	City         string `json:"city" bson:"city"`
	CityCode     string `json:"cityCode" bson:"cityCode"`
	District     string `json:"district" bson:"district"`
	AdCode       string `json:"adCode" bson:"adCode"`
	ProvinceCode string `json:"provinceCode" bson:"provinceCode"`
}

type AmapLocation struct {
	Status    string    `json:"status"`
	Regeocode Regeocode `json:"regeocode"`
}

type Regeocode struct {
	AddressComponent AddressComponent `json:"addresscomponent"`
}

type AddressComponent struct {
	Province     string `json:"province" bson:"province"`
	City         string `json:"city" bson:"city"`
	CityCode     string `json:"cityCode" bson:"cityCode"`
	District     string `json:"district" bson:"district"`
	AdCode       string `json:"adCode" bson:"adCode"`
	ProvinceCode string `json:"provinceCode" bson:"provinceCode"`
}
type LocationByRegeo struct {
	Longitude float64 `json:"longitude" form:"longitude"`
	Latitude  float64 `json:"latitude" form:"latitude"`
}

type LocationByGeo struct {
	Address string `json:"address"`
}
