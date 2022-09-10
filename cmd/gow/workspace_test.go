package main

import (
	"os"
	"testing"
)

func Test_workspace_struct_is_expected(t *testing.T) {
	t.Run("Constructed without params", func(t *testing.T) {
		// This is needed since working dir may vary depending on where the test is run
		// We expect the working dir to be the current dir + "/unnamed-project"
		expectWd, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}
		expectWd = expectWd + "/unnamed-project"
		var param1 string
		var param2 string
		w := workspace{}
		w.construct(param1, param2)

		got := w
		expected := workspace{workingDir: expectWd, packageDir: expectWd + "/pkg/", commandDir: expectWd + "/cmd/", mainDir: expectWd + "/cmd/main/", binariesDir: expectWd + "/bin/", mainFilePath: w.mainDir + "main.go", mainTestFilePath: w.mainDir + "main_test.go"}

		if got != expected {
			t.Errorf("GOT: %#q\nEXPECTED: %#q\n", got, expected)
		}
	})
	t.Run("Constructed with one parameter", func(t *testing.T) {
		w := workspace{}
		param1 := "test-project"
		var param2 string
		w.construct(param1, param2)
		got := w
		expected := workspace{workingDir: "./test-project", packageDir: "./test-project/pkg/", commandDir: "./test-project/cmd/", mainDir: "./test-project/cmd/main/", binariesDir: "./test-project/bin/", mainFilePath: "./test-project/cmd/main/main.go", mainTestFilePath: "./test-project/cmd/main/main_test.go"}

		if got != expected {
			t.Errorf("GOT: %q\nEXPECTED: %q\n", got, expected)
		}
	})
	t.Run("Constructed with both parameters", func(t *testing.T) {
		w := workspace{}
		param1 := "projectFoo"
		param2 := "appBar"
		w.construct(param1, param2)

		got := w
		// Assert that the constructed workspace has the expected working dir
		expected := workspace{workingDir: "./projectFoo", packageDir: "./projectFoo/pkg/", commandDir: "./projectFoo/cmd/", mainDir: "./projectFoo/cmd/appBar/", binariesDir: "./projectFoo/bin/", mainFilePath: "./projectFoo/cmd/appBar/main.go", mainTestFilePath: "./projectFoo/cmd/appBar/main_test.go"}

		if got != expected {
			t.Errorf("GOT: %q\nEXPECTED: %q\n", got, expected)
		}
	})
}
