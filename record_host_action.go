package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/rtucker-mozilla/go-infoblox"
	"os"
)

func record_host_get(hostname string, config Config) {
	ib := getInfobloxClient(config)
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

func record_host_create(hostname string, ipv4addrs string, configureForDHCP bool, mac string, view string, config Config) {
	ib := getInfobloxClient(config)

	addrs := []infoblox.HostIpv4Addr{
		infoblox.HostIpv4Addr{
			ConfigureForDHCP: configureForDHCP,
			MAC:              mac,
			Ipv4Addr:         ipv4addrs,
		},
	}

	host := infoblox.RecordHostObject{
		ConfigureForDNS: true,
		Ipv4Addrs:       addrs,
		Name:            hostname,
		View:            view,
	}
	resp, err := ib.CreateRecordHost(host)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	} else {
		fmt.Println("SUCCESS:", resp)
		os.Exit(0)
	}

}

func record_host_execute(action string, opts docopt.Opts, config Config) {
	hostname, _ := opts.String("<hostname>")
	if len(hostname) == 0 {
		fmt.Println("Hostname cannot be blank")
	}
	if action == "get" {
		record_host_get(hostname, config)
	} else if action == "create" {
		ipv4addrs, _ := opts.String("<ipv4addrs>")
		view, _ := opts.String("<view>")
		configureForDHCPVal, macErr := opts["--configure-for-dhcp"].(bool)
		mac, macErr := opts["--mac"].(string)
		if macErr != false {
			fmt.Println(macErr)
		}
		record_host_create(hostname, ipv4addrs, configureForDHCPVal, mac, view, config)
	}
}
