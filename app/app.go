package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI App"
	app.Usage = "Get Ips and serves in the network"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "get ips and address",
			Flags:  flags,
			Action: getIp,
		},
		{
			Name:   "server",
			Usage:  "get serve",
			Flags:  flags,
			Action: getServe,
		},
	}

	return app
}

func getIp(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getServe(c *cli.Context) {
	host := c.String("host")

	serves, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, serve := range serves {
		fmt.Println(serve)
	}
}
