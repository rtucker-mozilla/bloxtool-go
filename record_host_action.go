package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/rtucker-mozilla/go-infoblox"
	"os"
)

func record_host_get(hostname string, view string, config Config) {
	ib := getInfobloxClient(config)
	hosts, hostFoundErr := ib.FindRecordHost(hostname, view)
	if hostFoundErr != nil || len(hosts) == 0 {
		fmt.Println("Host Not Found")
		os.Exit(2)
	} else {
		for _, host := range hosts {
			fmt.Printf("Hostname: %s ref: %s", host.Name, host.Ref)
		}
	}

}

func record_host_delete(hostname string, view string, config Config) {
	ib := getInfobloxClient(config)
	hosts, hostFoundErr := ib.FindRecordHost(hostname, view)
	if hostFoundErr != nil {
		fmt.Println(hostFoundErr)
	}
	if len(hosts) == 0 {
		fmt.Printf("Error: Unable to find host %s in view: %s", hostname, view)
		os.Exit(2)
	}

	ref := hosts[0].Ref
	deleted, deletedErr := ib.RecordHost().Delete(ref)
	if deletedErr != nil {
		fmt.Println("Error:", deletedErr)
		os.Exit(2)
	}

	fmt.Println("Success:", deleted)

}
func RecordHostCreate(hostname string, ipv4addrs string, configureForDHCP bool, mac string, view string, config Config) {
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
		fmt.Println("Success:", resp)
		os.Exit(0)
	}

}

func record_host_execute(action string, opts docopt.Opts, config Config) {
	hostname, _ := opts.String("<hostname>")
	view, _ := opts.String("<view>")
	if len(hostname) == 0 {
		fmt.Println("Hostname cannot be blank")
	}
	if action == "get" {
		record_host_get(hostname, view, config)
	} else if action == "create" {
		ipv4addrs, _ := opts.String("<ipv4addrs>")
		configureForDHCPVal, _ := opts["--configure-for-dhcp"].(bool)
		mac, _ := opts["--mac"].(string)
		RecordHostCreate(hostname, ipv4addrs, configureForDHCPVal, mac, view, config)
	} else if action == "delete" {
		record_host_delete(hostname, view, config)
	}
}
