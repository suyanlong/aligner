package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

var (
	app     *cli.App
	comment string // 单行注释符号
	path    string // 格式化路径
	replace bool   // 是否替换
	ext     string // 扩展名
	isCheck bool
)

var langExt = map[string]string{
	".py":    "#", // python
	".sh":    "#", // shell
	".bash":  "#", // shell
	".yaml":  "#",
	".yml":   "#",
	".rb":    "#", // ruby
	".c":     "//",
	".cpp":   "//",
	".h":     "//",
	".java":  "//",
	".go":    "//",
	".rs":    "//",
	".js":    "//",
	".hpp":   "//", // C++
	".swift": "//",
	".m":     "//", // objectc
	".dart":  "//", // dart
	".kt":    "//", // kotlin
	".php":   "//",
	".cs":    "//",
	".scala": "//",
}

func init() {
	path, _ := os.Getwd()
	app = &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "comment",
				Aliases:     []string{"c"},
				Usage:       "annotation symbols",
				Destination: &comment,
			},
			&cli.StringFlag{
				Name:        "path",
				Aliases:     []string{"p"},
				Value:       path,
				Usage:       "alignment path",
				Destination: &path,
			},
			&cli.BoolFlag{
				Name:        "replace",
				Aliases:     []string{"r,w"},
				Value:       false,
				Usage:       "replace file",
				Destination: &replace,
			},
			&cli.StringFlag{
				Name:        "ext",
				Aliases:     []string{"e"},
				Usage:       "file extension",
				Destination: &ext,
			},
		},
		Commands: cli.Commands{
			&cli.Command{
				Name:        "comment",
				Usage:       "align multiple lines of single comments",
				UsageText:   "aligner comment " + PWD(),
				Description: "align multiple lines of single comments",
				Action:      commentAction,
			},
			&cli.Command{
				Name:        "check",
				Usage:       "aligner check whether the symbols are aligned",
				UsageText:   "aligner check " + PWD(),
				Description: "aligner check for CI/CD",
				Action:      checkAction,
			},
		},

		Name:    "aligner",
		Usage:   "A new CI/CD cli, align multiple lines of single comments or symbols",
		Version: "1.0.0",
	}
}

func commentAction(c *cli.Context) error {
	for _, val := range c.Args().Slice() {
		load(val)
	}
	return nil
}

func checkAction(c *cli.Context) error {
	isCheck = true
	for _, val := range c.Args().Slice() {
		load(val)
	}
	return nil
}

func PWD() string {
	path, err := os.Getwd()
	errorCheck(err)
	return path
}
