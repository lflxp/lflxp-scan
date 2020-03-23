package main

import (
	"flag"
	"fmt"

	"github.com/lflxp/lflxp-scan/pkg"
)

var ip string

func init() {
	flag.StringVar(&ip, "i", "192.168.1.1-255", "ip range,use -")
	flag.Parse()
}

func main() {
	pkg.Scan(fmt.Sprintf("scan %s", ip))
}
