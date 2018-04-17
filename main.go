package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"os"
	"os/user"
)

func printList(out []map[string]interface{}, err error) {
	e(err)
	for i, v := range out {
		fmt.Printf("[%d]\n", i)
		printObject(v, nil)
	}
}
func printObject(out map[string]interface{}, err error) {
	e(err)
	for k, v := range out {
		fmt.Printf("  %s: %q\n", k, v)
	}
	fmt.Printf("\n")
}

func printString(out string, err error) {
	e(err)
	fmt.Printf("  %q\n\n", out)
}

func printError(err error) {
	e(err)
	fmt.Printf("SUCCESS\n")
}

func e(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func main() {
	usage := `Usage:
    bloxtool record:host get <hostname> <view>
    bloxtool record:host create <hostname> <ipv4addrs> <view> [--mac=<mac>] [--configure-for-dhcp=<true>]
    bloxtool record:host delete <hostname> <view>`
	opts, _ := docopt.ParseDoc(usage)
	argv := os.Args[1:]
	//opts, err := parser.ParseArgs(usage, argv, "")
	// @TODO: add config file path option to bloxtool
	usr, _ := user.Current()
	configFilePath := fmt.Sprintf("%s/%s", usr.HomeDir, ".bloxtool.cfg")
	config, err := get_config(configFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if argv[0] == "record:host" {
		if len(argv) > 1 {
			RecordHostExecute(argv[1], opts, config)
		}
	}
}
