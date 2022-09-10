package main

import (
	"fmt"
	"log"
	"os"
)

type workspace struct {
	workingDir       string
	packageDir       string
	commandDir       string
	mainDir          string
	binariesDir      string
	mainFilePath     string
	mainTestFilePath string
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
	w.mainTestFilePath = w.mainDir + "main_test.go"
}

func (w *workspace) createWorkspace() {
	// Check if the workspace already exists
	if _, err := os.Stat(w.workingDir); os.IsNotExist(err) {
		os.MkdirAll(w.packageDir, os.ModePerm)
		os.MkdirAll(w.commandDir, os.ModePerm)
		os.MkdirAll(w.binariesDir, os.ModePerm)
		os.MkdirAll(w.mainDir, os.ModePerm)
	} else {
		log.Fatalf("Workspace %v already exists", w.workingDir)
	}
}

func (w *workspace) createFiles() {
	// For main.go
	mainFile, err := os.Create(w.mainFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer mainFile.Close()
	_, err = mainFile.WriteString(mainFileContents)
	if err != nil {
		log.Fatal(err)
	}
	// For main_test.go
	mainTestFile, err := os.Create(w.mainTestFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer mainTestFile.Close()
	_, err = mainTestFile.WriteString(mainTestFileContents)
	if err != nil {
		log.Fatal(err)
	}
}

func (w *workspace) create() {
	w.createWorkspace()
	w.createFiles()
	w.reminder()
}

func (w *workspace) reminder() {
	fmt.Println("✌\tRemember to run `go mod init [projectName]` in the workspace directory.")
}
