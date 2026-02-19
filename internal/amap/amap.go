package amap

import (
	"net/url"
	"rent/internal/config"
)

var mapUrl = url.Values{}

func InitAmap() {
	cfg := config.Cfg.Amap
	mapUrl.Add("key", cfg.Key)
	
}
