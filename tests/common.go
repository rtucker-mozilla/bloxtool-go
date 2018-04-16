package main

import (
	"github.com/rtucker-mozilla/go-infoblox"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
