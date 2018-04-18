package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/rtucker-mozilla/go-infoblox"
	"os"
)

func RecordCnameGet(alias string, view string, config Config) {
	ib := getInfobloxClient(config)
	cname, cnameFoundErr := ib.FindRecordCname(alias, view)
	if len(cname) == 0 || cnameFoundErr != nil {
		fmt.Println("cname Not Found")
		os.Exit(2)
	} else {
		for _, cname := range cname {
			fmt.Printf("cnamename: %s ref: %s", cname.Name, cname.Ref)
		}
	}

}

func RecordCnameDelete(alias string, view string, config Config) {
	ib := getInfobloxClient(config)
	cname, cnameFoundErr := ib.FindRecordCname(alias, view)
	if len(cname) == 0 || cnameFoundErr != nil {
		fmt.Println("cname Not Found")
		os.Exit(2)
	}

	ref := cname[0].Ref
	deleted, deletedErr := ib.RecordCname().Delete(ref)
	if deletedErr != nil {
		fmt.Println("Error:", deletedErr)
		os.Exit(2)
	}

	fmt.Println("Success:", deleted)

}

func RecordCnameCreate(alias string, cname string, view string, config Config) {
	ib := getInfobloxClient(config)

	cnameObj := infoblox.RecordCnameObject{
		Name:      alias,
		Canonical: cname,
		View:      view,
	}
	resp, err := ib.CreateRecordCname(cnameObj)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	} else {
		fmt.Println("Success:", resp)
		os.Exit(0)
	}

}

func RecordCnameExecute(action string, opts docopt.Opts, config Config) {
	alias, _ := opts.String("<alias>")
	view, _ := opts.String("<view>")
	if len(alias) == 0 {
		fmt.Println("alias cannot be blank")
	}
	if action == "get" {
		RecordCnameGet(alias, view, config)
	} else if action == "create" {
		cname, _ := opts.String("<cname>")
		RecordCnameCreate(alias, cname, view, config)
	} else if action == "delete" {
		RecordCnameDelete(alias, view, config)
	}
}
