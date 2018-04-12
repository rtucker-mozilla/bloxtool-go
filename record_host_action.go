package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/rtucker-mozilla/go-infoblox"
	"os"
)

func record_host_execute(action string, opts docopt.Opts, config Config) {
	if action == "get" {
		hostname, hostActionErr := opts.String("<hostname>")
		if hostActionErr != nil {
			fmt.Println("hostname required for get")
		}
		ib := infoblox.NewClient(config.infoblox_host, config.infoblox_username, config.infoblox_password, true, false)
		hosts, hostFoundErr := ib.FindRecordHost(hostname)
		if hostFoundErr != nil || len(hosts) == 0 {
			fmt.Println("Host Not Found")
			os.Exit(2)
		} else {
			for _, host := range hosts {
				fmt.Printf("Hostname: %s ref: %s", host.Name, host.Ref)
			}
		}
	}
}
