package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type workspace struct {
	workingDir   string
	packageDir   string
	commandDir   string
	mainDir      string
	binariesDir  string
	mainFilePath string
}

func (w *workspace) construct() {
	if w.workingDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		w.workingDir = wd + "/unnamed-project"
	}
	w.binariesDir = w.workingDir + "/bin/"
	w.commandDir = w.workingDir + "/cmd/"
	w.packageDir = w.workingDir + "/pkg/"
	if w.mainDir == "" {
		w.mainDir = w.commandDir + "main/"
	}
}

func (w *workspace) setMainFilePath() {
	w.mainFilePath = w.mainDir + "main.go"
}

func (w *workspace) createWorkspace() {
	os.MkdirAll(w.packageDir, os.ModePerm)
	os.MkdirAll(w.commandDir, os.ModePerm)
	os.MkdirAll(w.binariesDir, os.ModePerm)
	os.MkdirAll(w.mainDir, os.ModePerm)
	log.Printf("Created workspace at %s", w.workingDir)
}

func (w *workspace) createMainFile() {
	mainFile, err := os.Create(w.mainFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer mainFile.Close()
	_, err = mainFile.WriteString(mainFileContents)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created main.go file at %s", w.mainFilePath)
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
