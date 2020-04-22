package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/urfave/cli/v2"

	"github.com/suyanlong/aligner/version"
)

var (
	app     *cli.App
	comment string // 单行注释符号
	path    string // 格式化路径
	replace bool   // 是否替换
	ext     string // 扩展名
	ignore  string // 忽略指定的目录
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
			&cli.StringFlag{
				Name:        "ignore",
				Aliases:     []string{"i"},
				Value:       "",
				Usage:       "ignore file or direction",
				Destination: &ignore,
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
		Version: fmt.Sprintf(" aligner %s %s %s %s\n", version.Version, runtime.GOOS, runtime.GOARCH, version.BuildTime),
	}
}

func commentAction(c *cli.Context) error {
	commonAction(c)
	return nil
}

func commonAction(c *cli.Context) {
	if c.Args().Len() == 0 {
		load(PWD())
	}
	for _, val := range c.Args().Slice() {
		load(val)
	}
}

func checkAction(c *cli.Context) error {
	isCheck = true
	commonAction(c)
	return nil
}

func PWD() string {
	path, err := os.Getwd()
	errorCheck(err)
	return path
}
