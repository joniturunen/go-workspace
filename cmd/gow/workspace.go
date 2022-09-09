package main

import (
	"log"
	"os"
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
