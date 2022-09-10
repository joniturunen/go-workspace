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

func (w *workspace) construct(projectName string, appName string) {
	if projectName == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		w.workingDir = wd + unnamedProjectName
	} else {
		w.workingDir = "./" + projectName
	}
	w.binariesDir = w.workingDir + binDir
	w.commandDir = w.workingDir + cmdDir
	w.packageDir = w.workingDir + pkgDir
	if appName == "" {
		w.mainDir = w.commandDir + mainDirPostfix
	} else {
		w.mainDir = w.commandDir + appName + "/"
	}
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
