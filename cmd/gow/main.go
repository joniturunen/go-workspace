package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func reminder() {
	log.Println("Remember to run go mod init [projectName] in the workspace directory")
}

func main() {
	app := &cli.App{
		Name:  "gow",
		Usage: "gow [projectName]",
		Action: func(cCtx *cli.Context) error {
			w := workspace{}
			if cCtx.Args().Len() > 0 {
				w.workingDir = "./" + cCtx.Args().Get(0)
			}
			w.construct()
			if cCtx.Args().Len() == 2 {
				w.mainDir = w.commandDir + cCtx.Args().Get(1) + "/"
			}
			w.setMainFilePath()
			w.createWorkspace()
			w.createMainFile()
			reminder()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
