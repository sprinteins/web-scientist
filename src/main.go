package main

import (
	"flag"
	"fmt"

	server "github.com/sprinteins/web-scientist/server"
	"github.com/sprinteins/web-scientist/utils"
)

func main() {
	var host = flag.String("host", "localhost", "The Host")
	var port = flag.String("port", "7654", "The port")
	var reference = flag.String("reference", "", "The reference service")
	var experiment = flag.String("experiment", "", "The experiment service")

	utils.TellMeMore()

	flag.Parse()
	fmt.Printf("listening on http://%s:%s\n", *host, *port)
	// server(*host, *port)
	fmt.Printf("server started")
	// server.Start(*host, *port)
	var scientist = server.New(*host, *port)
	scientist.SetReference(*reference)
	scientist.SetExperiment(*experiment)
	scientist.Start()

}
