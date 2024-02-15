package main

import (
	"fmt"
	"log"
	"os"

	"git.garena.com/bpf/tc"
	cli "github.com/urfave/cli/v2"
)

// 初始化+调用入口

func main() {
	// new tc
	tc := tc.NewTCclient()

	app := &cli.App{
		Name:  "tc",
		Usage: "tencent cloud instance action !",
		Commands: []*cli.Command{
			{
				Name:    "destory",
				Aliases: []string{"d"},
				Usage:   "stop an instance with special tag",
				Action: func(cCtx *cli.Context) error {
					key := cCtx.Args().Get(0)
					value := cCtx.Args().Get(1)
					fmt.Println("return an instance with instance key:value: ", key, value)
					tc.Destory(key, value)
					return nil
				},
			},
			{
				Name:    "start",
				Aliases: []string{"s", "r"},
				Usage:   "start an instance with template",
				Action: func(cCtx *cli.Context) error {
					tc.Run()
					fmt.Println("start an instance with instance tag: ")
					return nil
				},
			},
			{
				Name:    "query",
				Aliases: []string{"q"},
				Usage:   "query instances",
				Action: func(cCtx *cli.Context) error {
					key := cCtx.Args().Get(0)
					value := cCtx.Args().Get(1)
					tc.Describe(key, value)
					fmt.Println("query instance with speical tag: ", key, value)
					return nil
				},
			},
		},
		Action: func(*cli.Context) error {
			str := `
				tc 目前仅支持启动实例和销毁指定实例
				tc start 启动一个默认的实力，该实例具有特定的模板
				tc destory tag 销毁一个实例，如果给定tag,则销毁对应的tag 实例，否则销毁所有实例
			`
			fmt.Println(str)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
