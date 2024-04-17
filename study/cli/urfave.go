package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	app := cli.App{
		Name:  "testa",
		Usage: "asdfasfasfasf",
		Action: func(context *cli.Context) error {
			file := context.String("file")
			fmt.Printf("main run:%s\n", file)
			return nil
		},
		After: func(context *cli.Context) error {
			fmt.Printf("run after function\n")
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Value: "55",
				Usage: "file path",
				Action: func(context *cli.Context, s string) error {
					fmt.Printf("run %s flag\n", "file")
					return nil
				},
			},
		},
		Commands: []*cli.Command{
			{
				Name: "log1",
				Action: func(context *cli.Context) error {
					fmt.Printf("run %s command\n", "log1")
					return fmt.Errorf("test error a")
				},
			},
			{
				Name: "log2",
				Subcommands: []*cli.Command{
					{
						Name: "l1",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:  "a",
								Usage: "aa",
							},
						},
						Action: func(context *cli.Context) error {
							fmt.Printf("run %s command,%s subcommand,%v flag\n", "log1", "l1", context.Bool("a"))
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("happen error,%v", err)
	}

}
