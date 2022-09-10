package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gow",
		Usage: "gow [projectName]",
		Action: func(cCtx *cli.Context) error {
			w := workspace{}
			if cCtx.Args().Len() == 1 {
				p1 := cCtx.Args().Get(0)
				w.construct(p1, "")
			}
			if cCtx.Args().Len() == 2 {
				p1 := cCtx.Args().Get(0)
				p2 := cCtx.Args().Get(1)
				w.construct(p1, p2)
			}
			// fmt.Printf("%#v\n", w)
			w.create()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
