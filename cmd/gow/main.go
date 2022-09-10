package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	w := workspace{}
	var projectName string
	var appName string
	app := &cli.App{
		Name:  "gow",
		Usage: "gow [projectName]",
		Action: func(cCtx *cli.Context) error {
			if cCtx.Args().Len() == 1 {
				projectName = cCtx.Args().Get(0)
				w.construct(projectName, "")
			} else if cCtx.Args().Len() == 2 {
				projectName = cCtx.Args().Get(0)
				appName = cCtx.Args().Get(1)
				w.construct(projectName, appName)
			} else {
				w.construct(projectName, appName)
			}
			w.create()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
