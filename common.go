package main

import (
	"github.com/rtucker-mozilla/go-infoblox"
)

func getInfobloxClient(config Config) *infoblox.Client {
	ib := infoblox.NewClient(config.infoblox_host, config.infoblox_username, config.infoblox_password, false, false)
	return ib
}
