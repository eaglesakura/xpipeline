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
	app.Version = "0.0"

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
