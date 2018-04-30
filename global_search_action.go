package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"os"
)

func GlobalSearchExecute(term string, opts docopt.Opts, config Config) {
	if len(term) == 0 {
		fmt.Println("term cannot be blank")
		os.Exit(2)
	}
	ib := getInfobloxClient(config)
	objtype, _ := opts["--objtype"].(string)
	s_return, err := ib.Search(term, objtype)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}
	fmt.Printf(s_return)
}
