package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/cheggaaa/pb/v3"
	"github.com/huandu/xstrings"
)

var (
	enable  = "aligner enable"
	disable = "aligner disable"
	todo    = "todo"
)

type Context struct {
	path     string // file path
	contexts []string
	isEnable bool // default true
}

func New(path string) *Context {
	return &Context{
		isEnable: true,
		path:     path,
	}
}

func (c *Context) FormatFile() {
	f, err := os.Open(c.path)
	errorCheck(err)
	defer f.Close()
	reader := bufio.NewReader(f)
	var section []string
	maxDistance := 0
	transfer := func() {
		for _, val := range section {
			lastIndex := strings.LastIndex(val, comment)
			diff := maxDistance - lastIndex
			if isCheck && diff > 0 {
				info := "file : " + c.path + ", context: "
				fmt.Println(info, section)
				os.Exit(1)
			}
			blank := addBlankString(diff)
			newLine := xstrings.Insert(val, blank, lastIndex)
			c.contexts = append(c.contexts, newLine)
		}
		section = section[:0]
		maxDistance = 0
	}

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err != io.EOF {
			panic(err)
		}
		l := string(line)
		if err == io.EOF {
			transfer()
			break
		}
		c.switched(l)
		if c.isSkip(l) {
			transfer()
			c.contexts = append(c.contexts, l)
			continue
		}

		if !strings.Contains(l, comment) {
			transfer()
			c.contexts = append(c.contexts, l)
		} else {
			lastIndex := strings.LastIndex(l, comment)
			if lastIndex > maxDistance {
				maxDistance = lastIndex
			}
			section = append(section, l)
		}
	}
	c.writer()
}

func (c *Context) writer() {
	if replace {
		// TODO
		out, err := os.Create(c.path)
		defer out.Close()
		errorCheck(err)
		err = writer(out, c.contexts)
		errorCheck(err)

	} else {
		fmt.Println("-----------" + c.path + "-----------")
		err := writer(os.Stdout, c.contexts)
		errorCheck(err)
		fmt.Println("-----------" + c.path + "-----------")
	}
}

func (c *Context) switched(line string) {
	if strings.Contains(line, enable) {
		c.isEnable = true
	}
	if strings.Contains(line, disable) {
		c.isEnable = false
	}
}

func (c *Context) isSkip(line string) bool {
	if !c.isEnable {
		return true
	}

	if strings.Contains(line, todo) || strings.Contains(line, strings.ToUpper(todo)) {
		return true
	}

	l := strings.TrimSpace(line)
	return strings.Index(l, comment) <= 1
}

func IsIgnore(path string) bool {
	if ignore == "" {
		return false
	}
	pi, err := filepath.Abs(ignore)
	errorCheck(err)
	p, err := filepath.Rel(PWD(), pi)
	errorCheck(err)
	return strings.Contains(path, p)
}

func IsDir(p string) bool {
	s, err := os.Stat(p)
	errorCheck(err)
	return s.IsDir()
}

func IsDotFile(p string) bool {
	return strings.Index(filepath.Base(p), ".") == 0
}

func IsFormatFile(p string) bool {
	e := filepath.Ext(p)
	if e != "" && e == ext {
		if comment == "" {
			c, is := langExt[e]
			if !is {
				return is
			}
			comment = c // TODO
		}
		return true
	}

	if ext == "" && comment != "" && e != "" {
		return true
	}
	return false
}

func load(rootPath string) {

	var count int64 = 0
	// create and start new bar
	bar := pb.StartNew(int(count))
	// finish bar
	defer bar.Finish()

	err := filepath.Walk(
		rootPath,
		func(path string, info os.FileInfo, err error) error {
			if IsIgnore(path) {
				return nil
			}
			if info.IsDir() {
				return nil
			}
			if IsDotFile(path) {
				return nil
			}
			if IsFormatFile(path) {
				New(path).FormatFile()
				count++
				bar.SetCurrent(count - 2)
				bar.SetTotal(count)
				bar.Increment()
			}
			return err
		},
	)
	errorCheck(err)

}

func addBlankString(num int) string {
	blankStr := ""
	for i := 0; i < num; i++ {
		blankStr += " "
	}
	return blankStr
}

func TmpDir() string {
	return filepath.Join(PWD(), "tmp")
}

func writer(w io.Writer, contexts []string) error {
	bw := bufio.NewWriter(w)
	for _, val := range contexts {
		// lineStr := fmt.Sprintf("%s", val)
		_, err := fmt.Fprintln(bw, val)
		if err != nil {
			return err
		}
	}
	return bw.Flush()
}
