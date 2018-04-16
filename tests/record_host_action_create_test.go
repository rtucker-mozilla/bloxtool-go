package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRecordHostFound2(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/wapi/v1.4.1/record:host", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("record_host/record_host_found.json"))
	})
	hostname := "foo.domain.com"
	hosts, hostFoundErr := client.FindRecordHost(hostname)
	if hostFoundErr != nil {
		t.Error("Host should have been found")
	}
	if len(hosts) != 1 {
		t.Error("Improper Length of Hosts")
	}
	host := hosts[0]
	if host.Name != "foo.domain.com" {
		t.Error("host.Name not set properly")
	}
	if host.View != "Private" {
		t.Error("host.View not set properly")
	}
	if host.Ipv4Addrs[0].Ipv4Addr != "192.168.0.1" {
		t.Error("host.Ipv4Addrs.Ipv4Addr not set properly")
	}

}

func TestRecordHostNotFound2(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/wapi/v1.4.1/record:host", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("record_host/record_host_not_found.json"))
	})
	hostname := "foo.domain.com"
	hosts, hostFoundErr := client.FindRecordHost(hostname)
	if hostFoundErr != nil {
		t.Error("hostFoundErr should be nil")
	}
	if len(hosts) != 0 {
		t.Error("Improper Length of Hosts")
	}
}
