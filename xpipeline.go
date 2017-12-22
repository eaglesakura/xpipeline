package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "xpipeline"
	app.Description = "Mobile app build pipeline utils"
	app.Version = "0.1"

	//app.Flags = []cli.Flag{
	//	cli.StringFlag{
	//		Name:  "config",
	//		Usage: "`project.[yaml|yml]` path.",
	//	},
	//}
	app.Commands = []cli.Command{
		{
			Name: "mipmap",
			Action: func(ctx *cli.Context) {
				if task, err := newMipmapTask(ctx); err != nil {
					fmt.Errorf("%v\n", err)
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
		{
			Name: "yaml2json",
			Action: func(context *cli.Context) {
				if task, err := newYaml2JsonTask(context); err != nil {
					fmt.Errorf("%v\n", err)
				} else {
					task.Execute()
				}
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input",
					Usage: "path/to/input.[yaml|yml]",
				},
				cli.StringFlag{
					Name:  "output",
					Usage: "path/to/output.json",
				},
			},
		},
		{
			Name: "docker-path",
			Action: func(context *cli.Context) {
				if task, err := newDockerPathTask(context); err != nil {
					fmt.Errorf("%v\n", err)
				} else {
					task.Execute()
				}
			},
		},
	}
	app.Run(os.Args)
}
