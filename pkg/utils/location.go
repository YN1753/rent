package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"rent/internal/config"
	"rent/internal/model"
	"strconv"
	"strings"

	"github.com/valyala/fastjson"
)

func GetLocationByRegeo(loc model.LocationByRegeo) (model.Location, error) {
	cfg := config.Cfg
	res, err := http.Get("https://restapi.amap.com/v3/geocode/regeo" + "?" + "key=" + cfg.Amap.Key + fmt.Sprintf("&location=%f,%f", loc.Longitude, loc.Latitude))
	if err != nil {
		log.Println(err)
		return model.Location{}, errors.New("获取地理位置失败")
	}
	result, err := io.ReadAll(res.Body)
	if err != nil {
		return model.Location{}, errors.New("获取地理位置失败")
	}
	fmt.Println(string(result))
	var Amplocation model.AmapLocation
	json.Unmarshal(result, &Amplocation)
	var location model.Location
	fmt.Println(Amplocation)
	location.AdCode = Amplocation.Regeocode.AddressComponent.AdCode
	location.City = Amplocation.Regeocode.AddressComponent.City
	location.CityCode = Amplocation.Regeocode.AddressComponent.CityCode
	location.District = Amplocation.Regeocode.AddressComponent.District
	location.Province = Amplocation.Regeocode.AddressComponent.Province
	location.ProvinceCode = Amplocation.Regeocode.AddressComponent.ProvinceCode
	return location, nil
}

func GetLocationByGeo(address string) (model.LocationByRegeo, error) {
	cfg := config.Cfg
	params := url.Values{}
	var p fastjson.Parser
	params.Set("key", cfg.Amap.Key)
	params.Set("address", address)
	res, err := http.Get("https://restapi.amap.com/v3/geocode/geo" + "?" + params.Encode())
	if err != nil {
		return model.LocationByRegeo{}, errors.New("获取地址失败")
	}
	response, err := io.ReadAll(res.Body)
	if err != nil {
		return model.LocationByRegeo{}, errors.New("获取地址失败")
	}
	result, err := p.Parse(string(response))
	if err != nil {
		return model.LocationByRegeo{}, errors.New("json解析失败")
	}
	var location model.LocationByRegeo
	location_pair := result.GetArray("geocodes")[0].Get("location").String()
	location_pair = strings.ReplaceAll(location_pair, "\"", "")
	array := strings.Split(location_pair, ",")
	location.Latitude, err = strconv.ParseFloat(array[0], 64)
	if err != nil {
		return model.LocationByRegeo{}, errors.New("获取地址失败")
	}
	location.Longitude, err = strconv.ParseFloat(array[1], 64)
	if err != nil {
		return model.LocationByRegeo{}, errors.New("获取地址失败")
	}
	return location, nil
}
