package main

const mainFileContents string = `package main

import (
	"log"
)

func main() {
	log.Println("Im waiting to be implemented")
} 
`

const mainTestFileContents string = `package main

import (
	"testing"
)

func Test_example(t *testing.T) {
	t.Run("Constructed with one parameter", func(t *testing.T) {
		got := "Assertion"
		expected := "Expected result"

		if got != expected {
			t.Errorf("GOT: %q\nEXPECTED: %q\n", got, expected)
		}
	})
}

`

const unnamedProjectName string = "/unnamed-project"
const binDir = "/bin/"
const cmdDir = "/cmd/"
const pkgDir = "/pkg/"
const mainDirPostfix = "main/"
