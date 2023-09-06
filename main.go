package main

import (
	"log"
	"os"

	dirInit "github.com/nelsen129/tskr/internal/init"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "initialize the directory structure",
				Action: func(*cli.Context) error {
					dirInit.Init(".")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
