package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/huandu/xstrings"
	"github.com/urfave/cli/v2"
)

var comment string
var path string
var replace bool

func init() {
	path, _ := os.Getwd()
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "comment",
				Aliases:     []string{"c"},
				Value:       "#",
				Usage:       "annotation symbols",
				Destination: &comment,
			},
			&cli.StringFlag{
				Name:        "path",
				Aliases:     []string{"p"},
				Value:       path,
				Usage:       "annotation symbols",
				Destination: &path,
			},
			&cli.BoolFlag{
				Name:        "replace",
				Aliases:     []string{"r,w"},
				Value:       true,
				Usage:       "replace file",
				Destination: &replace,
			},
		},
		Name:    "aligner",
		Usage:   "aliger fmt ./...",
		Version: "1.0.0",
		// Action: func(c *cli.Context) error {
		// 	return nil
		// },
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("------")
	format("./p")
}

func addBlankString(num int) string {
	blankStr := ""
	for i := 0; i < num; i++ {
		blankStr += " "
	}
	return blankStr
}

func format(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	var contexts []string
	var section []string
	var maxDistance = 0
	for {
		line, _, err := reader.ReadLine()
		l := string(line)
		if !strings.Contains(l, comment) || err == io.EOF {
			for _, val := range section {
				lastIndex := strings.LastIndex(val, comment)
				blank := addBlankString(maxDistance - lastIndex)
				newLine := xstrings.Insert(val, blank, lastIndex)
				contexts = append(contexts, newLine)
			}
			contexts = append(contexts, l)
			section = section[:0]
			maxDistance = 0
			if err == io.EOF {
				break
			}
		} else {
			lastIndex := strings.LastIndex(l, comment)
			if lastIndex > maxDistance {
				maxDistance = lastIndex
			}
			section = append(section, l)
		}
	}

	newFile := path + ".ok"
	wf, err := os.Create(newFile)
	if err != nil {
		panic(err)
	}
	defer wf.Close()
	w := bufio.NewWriter(wf)
	for _, val := range contexts {
		lineStr := fmt.Sprintf("%s", val)
		fmt.Fprintln(w, lineStr)
	}
	w.Flush()
}
