package main

import (
	"fmt"
	"github.com/hzxiao/neo-thinsdk-go/neo"
	"github.com/hzxiao/neo-thinsdk-go/utils"
	"github.com/spf13/pflag"
	"os"
)

const version = "0.0.1"

var (
	ver  = pflag.BoolP("version", "v", false, "print version")
	help = pflag.BoolP("help", "h", false, "show usage of neoconv command")
	sh2a = pflag.StringP("sh2a", "", "", "convert script hash to address")
)

func main() {
	pflag.Parse()

	if *ver {
		fmt.Printf("neotx%v\n", version)
		os.Exit(0)
	}
	if *help {
		pflag.Usage()
		os.Exit(0)
	}

	if *sh2a != "" {
		b, _ := utils.ToBytes(*sh2a)
		address, _ := neo.GetAddressFromScriptHash(b)
		fmt.Println(address)
	}
}
