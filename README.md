# Gow - Create Golang workspaces

> Creates a simple workspace for Go projects

## Description

`gow` creates an empty [workspace structure](https://www.developer.com/languages/go-project-layout/) in a sub-directory of the current directory.

Application will also create the `main.go` file in the `cmd/main` directory.


```text
go-project-name
|
|\_cmd
|   \_main
|        \_main.go
|\_pkg
|
\_go.mod
\_go.sum
```

The `main.go` file will contain the following code:

```go
package main

func main() {
    // implementation goes here
}
```

## Usage

Call the program `gow` with the name of the project as the first argument.

```bash
$ gow go-project-name
# OR
$ gow go-project-name application-name
```

Latter will create the `application-name` directory in the `cmd` directory in place of `main`.

## Building

To Build the gow application, run the following command:

```bash
$ go build -o bin/gow ./cmd/gow
```

## Todo

- Add flags to give parameters to the application (instead of possitional arguments)
- Add some other directory structure flags, e.g. for creating Gorilla Mux projects or Gin projects.