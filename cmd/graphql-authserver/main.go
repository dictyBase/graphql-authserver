package main

import (
	"log"
	"os"

	"github.com/dictyBase/graphql-authserver/internal/app/server"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "graphql-authserver"
	app.Usage = "cli for graphql-authserver"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log-format",
			Usage: "format of the logging out, either of json or text.",
			Value: "json",
		},
		cli.StringFlag{
			Name:  "log-level",
			Usage: "log level for the application",
			Value: "error",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "start-server",
			Usage:  "starts the graphql authserver",
			Action: server.RunServer,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "public-key, pkey",
					Usage:    "public key file for verifying jwt",
					EnvVar:   "JWT_PUBLIC_KEY",
					Required: true,
				},
				cli.StringFlag{
					Name:     "private-key, prkey",
					Usage:    "private key file for signing jwt",
					EnvVar:   "JWT_PRIVATE_KEY",
					Required: true,
				},
				cli.IntFlag{
					Name:  "port, p",
					Usage: "server port",
					Value: 9099,
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("error in running command %s", err)
	}
}
