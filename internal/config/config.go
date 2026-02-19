package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server  ServerConfig
	Mysql   MysqlConfig
	JWT     JWTConfig
	Redis   RedisConfig
	Email   EamilConfig
	Amap    AmapConfig
	MongoDB MongoDbConfig
}

type ServerConfig struct {
	Port string
	Mode string
}

type MysqlConfig struct {
	Driver string
	Source string
}

type RedisConfig struct {
	Addr     string
	Password string
	Db       int
}

type EamilConfig struct {
	Port     int
	Host     string
	Username string
	Password string
}
type MongoConfig struct {
	Uri      string
	Database string
}

type JWTConfig struct {
	Secret string
}

type AmapConfig struct {
	Key string
}

type MongoDbConfig struct {
	Url      string
	Port     string
	Database string
}

var Cfg Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(viper.Get("server.port"))
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		return
	}
}
