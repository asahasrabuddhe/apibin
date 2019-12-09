package main

import (
	"github.com/urfave/cli/v2"
	"go.ajitem.com/go-httpbin/apibin"
	"log"
	"os"
)

var Version string

func main() {
	app := cli.NewApp()

	app.Name = "apibin"
	app.Usage = "A simple service for testing API requests and responses"
	app.Version = Version

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "hostname",
			Aliases: []string{"n"},
		},
		&cli.IntFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Value:   8080,
		},
	}

	app.Action = func(context *cli.Context) error {
		hostname := context.String("hostname")
		port := context.Int("port")

		log.Printf("starting server on %s:%d\n", hostname, port)

		return apibin.Serve(hostname, port)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
