package main

import (
	"flag"

	"github.com/jadidbourbaki/gofarm/service"
)

// port is the port the service runs on
var port int

func parseFlags() {
	flag.IntVar(&port, "p", 443, "port to run service on")
	flag.Parse()

}

func main() {
	parseFlags()

	service := service.New()
	defer service.Shutdown()
	service.Run(port, true)
}
