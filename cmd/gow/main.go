package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type workspace struct {
	// The wd directory of the workspace.
	wd string
	// The directory containing the Go packages.
	pkg string
	// The directory containing the Go source files.
	cmd string
	// The directory containing the main.go file.
	main string
	// Location of main.go file.
	mainFile string
}

func (w *workspace) construct() {
	if w.wd == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		w.wd = wd + "/unnamed-project"
	}
	w.cmd = w.wd + "/cmd/"
	w.pkg = w.wd + "/pkg/"
	if w.main == "" {
		w.main = w.cmd + "main/"
	}
}

func (w *workspace) mainFilePath() {
	w.mainFile = w.main + "main.go"
}

func (w *workspace) createWorkspace() {
	os.MkdirAll(w.pkg, os.ModePerm)
	os.MkdirAll(w.cmd, os.ModePerm)
	os.MkdirAll(w.main, os.ModePerm)
	log.Printf("Created workspace at %s", w.wd)
}

func (w *workspace) createMainFile() {
	mainFile, err := os.Create(w.mainFile)
	if err != nil {
		log.Fatal(err)
	}
	defer mainFile.Close()
	_, err = mainFile.WriteString(mainFileContents)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created main.go file at %s", w.mainFile)
}

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
				w.wd = "./" + cCtx.Args().Get(0)
			}
			w.construct()
			if cCtx.Args().Len() == 2 {
				w.main = w.cmd + cCtx.Args().Get(1) + "/"
			}
			w.mainFilePath()
			w.createWorkspace()
			w.createMainFile()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
