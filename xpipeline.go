package main

import (
	"./mipmap"
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "xpipeline"
	app.Version = "0.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Usage: "`project.[yaml|yml]` path.",
		},
	}
	app.Commands = []cli.Command{
		{
			Name: "mipmap",
			Action: func(ctx *cli.Context) {
				if task, err := mipmap.NewInstance(ctx); err != nil {
					fmt.Printf("task error[%v]", err)
				} else {
					task.Execute()
				}

			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config",
					Usage: "mipmap.[yaml|yml] path",
				},
			},
		},
	}
	app.Run(os.Args)
}
