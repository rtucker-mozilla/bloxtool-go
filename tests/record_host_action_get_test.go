package main

import (
	"fmt"
	"github.com/rtucker-mozilla/go-infoblox"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux          *http.ServeMux
	server       *httptest.Server
	client       *infoblox.Client
	WAPI_VERSION = "1.4.1"
	BASE_PATH    = "/wapi/v" + WAPI_VERSION
	FULL_PATH    = BASE_PATH + "/" + "record:host"
)

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	username := "username"
	password := "password"
	client = infoblox.NewClient(server.URL, username, password, true, false)
	return func() {
		server.Close()
	}
}

func TestRecordHostFound(t *testing.T) {
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

func TestRecordHostNotFound(t *testing.T) {
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
