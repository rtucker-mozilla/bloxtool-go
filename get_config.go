package main

import (
	"errors"
	"gopkg.in/ini.v1"
)

type Config struct {
	infoblox_host     string
	infoblox_username string
	infoblox_password string
}

func get_config(config_path string) (Config, error) {
	cfg, err := ini.Load(config_path)
	if err != nil {
		return Config{}, errors.New("Unable to read config file")
	}
	infoblox_host := cfg.Section("InfoBlox").Key("host").String()
	if infoblox_host == "" {
		return Config{}, errors.New("host not found in config file")
	}
	infoblox_username := cfg.Section("InfoBlox").Key("username").String()
	if infoblox_username == "" {
		return Config{}, errors.New("username not found in config file")
	}
	infoblox_password := cfg.Section("InfoBlox").Key("password").String()
	if infoblox_password == "" {
		return Config{}, errors.New("password not found in config file")
	}

	config := Config{
		infoblox_host:     infoblox_host,
		infoblox_username: infoblox_username,
		infoblox_password: infoblox_password,
	}

	return config, nil

}
