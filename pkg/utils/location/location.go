package location

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fastjson"
	"io"
	"log"
	"net/http"
	"net/url"
	"rent/internal/config"
	"rent/internal/model"
	"strconv"
	"strings"
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
	var AmpLocation model.AmapLocation
	if err := json.Unmarshal(result, &AmpLocation); err != nil {
		return model.Location{}, err
	}
	var location model.Location
	fmt.Println(AmpLocation)
	location.AdCode = AmpLocation.Regeocode.AddressComponent.AdCode
	location.City = AmpLocation.Regeocode.AddressComponent.City
	location.CityCode = AmpLocation.Regeocode.AddressComponent.CityCode
	location.District = AmpLocation.Regeocode.AddressComponent.District
	location.Province = AmpLocation.Regeocode.AddressComponent.Province
	location.ProvinceCode = AmpLocation.Regeocode.AddressComponent.ProvinceCode
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
	locationPair := result.GetArray("geocodes")[0].Get("location").String()
	locationPair = strings.ReplaceAll(locationPair, "\"", "")
	array := strings.Split(locationPair, ",")
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

func GetLocationTips(keywords string) (model.LocationTipsRes, error) {
	cfg := config.Cfg
	params := url.Values{}
	var p fastjson.Parser
	var locationTips model.LocationTipsRes
	params.Set("key", cfg.Amap.Key)
	params.Set("keywords", keywords)
	url_ := "https://restapi.amap.com/v3/assistant/inputtips"
	response, err := http.Get(url_ + "?" + params.Encode())
	if err != nil {
		return model.LocationTipsRes{}, errors.New("获取地址失败")
	}
	resp, err := io.ReadAll(response.Body)
	if err != nil {
		return model.LocationTipsRes{}, errors.New("获取地址失败")
	}
	res, err := p.Parse(string(resp))
	if res == nil {
		return model.LocationTipsRes{}, err
	}
	for _, v := range res.GetArray("tips") {
		locationTips.Locations = append(locationTips.Locations, strings.ReplaceAll(v.Get("name").String(), "\"", ""))
	}
	return locationTips, nil
}
func GetLocationByPOI(keywords string) ([]model.LocationPOIRes, error) {
	var cfg = config.Cfg
	var LocationsList = make([]model.LocationPOIRes, 0)
	url_ := "https://restapi.amap.com/v3/place/text"
	var param = url.Values{}
	param.Set("key", cfg.Amap.Key)
	param.Set("keywords", keywords)
	param.Set("extensions", "all")
	response, err := http.Get(url_ + "?" + param.Encode())
	if err != nil {
		return nil, err
	}
	resp, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var p fastjson.Parser
	result, err := p.Parse(string(resp))
	pois := result.GetArray("pois")
	for i, v := range pois {
		var temp model.LocationPOIRes
		temp.City = strings.ReplaceAll(v.Get("cityname").String(), "\"", "")
		temp.AdCode = strings.ReplaceAll(v.Get("adcode").String(), "\"", "")
		temp.CityCode = strings.ReplaceAll(v.Get("citycode").String(), "\"", "")
		temp.Province = strings.ReplaceAll(v.Get("pname").String(), "\"", "")
		temp.ProvinceCode = strings.ReplaceAll(v.Get("pcode").String(), "\"", "")
		temp.Name = strings.ReplaceAll(v.Get("name").String(), "\"", "")
		locationPair := strings.ReplaceAll(v.Get("location").String(), "\"", "")
		locationPair = strings.ReplaceAll(locationPair, "\"", "")
		array := strings.Split(locationPair, ",")
		temp.Latitude, err = strconv.ParseFloat(array[0], 64)
		if err != nil {
			return nil, err
		}
		temp.Longitude, err = strconv.ParseFloat(array[1], 64)
		if err != nil {
			return nil, err
		}
		temp.District = strings.ReplaceAll(v.Get("adname").String(), "\"", "")
		if i > 5 {
			break
		}
		LocationsList = append(LocationsList, temp)
	}
	return LocationsList, nil
}
