package config

import (
	"encoding/json"
	"errors"
	"os"
)

var (
	ErrInvalidConfig = errors.New("invalid service config")
	ErrConfigIssues  = errors.New("not Able to open config file")
)

type Config struct {
	ServiceName string `json:"service_name"`

	RedisAddr string `json:"redis_addr"`
	HttpPort  int    `json:"http_port"`

	TokenKey Token `json:"token_key"`
}

type Token struct {
	AccessSecret  string `json:"access_secret"`
	RefreshSecret string `json:"refresh_secret"`
	AtExpires     int    `json:"at_expires"`
	RtExpires     int    `json:"rt_expires"`
}

func NewConfig(Isdefault bool) (c Config, err error) {

	//common\config\cfg.json
	configFile, err := os.Open("common/config/cfg.json")

	if err != nil {
		panic(err)

	}
	defer configFile.Close()

	err = json.NewDecoder(configFile).Decode(&c)
	

	if err != nil {
		panic("unable to read config file")
	}

	if !c.isValid() {
		err = ErrInvalidConfig
		return
	}

	return
}

func (c *Config) isValid() bool {
	if len(c.ServiceName) < 1 {
		return false
	}

	if len(c.RedisAddr) < 1 {
		return false
	}

	if c.HttpPort <= 0 {
		return false
	}

	if len(c.TokenKey.AccessSecret) < 1 {
		return false
	}

	if len(c.TokenKey.RefreshSecret) < 1 {
		return false
	}

	if c.TokenKey.AtExpires <= 0 {
		return false
	}

	if c.TokenKey.RtExpires <= 0 {
		return false
	}

	return true
}
