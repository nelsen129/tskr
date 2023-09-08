package main

import (
	"log"
	"os"

	dirInit "github.com/nelsen129/tskr/internal/init"
	"github.com/urfave/cli/v2"
)

func run(args []string) {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "initialize the directory structure",
				Action: func(*cli.Context) error {
					return dirInit.Init(".")
				},
			},
		},
	}

	if err := app.Run(args); err != nil {
		log.Printf("[ERROR] %v", err)
	}
}

func main() {
	run(os.Args)
}
