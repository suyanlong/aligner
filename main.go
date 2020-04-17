package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/huandu/xstrings"
)

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	errorCheck(app.Run(os.Args))
}

func format(p string) {
	if IsDir(p) {
		load(p)
	} else {
		formatFile(p)
	}
}

func IsDir(p string) bool {
	s, err := os.Stat(p)
	errorCheck(err)
	return s.IsDir()
}

func IsDotFile(p string) bool {
	if strings.Index(filepath.Base(p), ".") == 0 {
		return true
	}
	return false
}

func IsFormatFile(p string) bool {
	e := filepath.Ext(p)
	if ext != "" && ext == e {

		return true
	}
	if e != "" {
		c, ok := langExt[e]
		comment = c // TODO
		return ok
	}
	//panic("file extension not exist!")
	return false
}

//func selfAdjust(file string)  {
//	c, ok := langExt[e]
//	comment = c // TODO
//	return ok
//}

func load(rootPath string) {
	err := filepath.Walk(
		rootPath,
		func(path string, info os.FileInfo, err error) error {
			fmt.Println(path)
			if info.IsDir() {
				return nil
			}
			if IsDotFile(path) {
				return nil
			}
			if IsFormatFile(path) {
				formatFile(path)
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

func formatFile(path string) {
	f, err := os.Open(path)
	errorCheck(err)
	defer f.Close()
	reader := bufio.NewReader(f)
	var contexts []string
	var section []string
	maxDistance := 0
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

	if replace {
		// TODO
		out, err := os.Create(path)
		defer out.Close()
		errorCheck(err)
		err = writer(out, contexts)
		errorCheck(err)

	} else {
		fmt.Println("-----------" + path + "-----------")
		err = writer(os.Stdout, contexts)
		errorCheck(err)
		fmt.Println("-----------" + path + "-----------")
	}
}

func TmpDir() string {
	return filepath.Join(PWD(), "tmp")
}

func writer(w io.Writer, contexts []string) error {
	bw := bufio.NewWriter(w)
	for _, val := range contexts {
		lineStr := fmt.Sprintf("%s", val)
		_, err := fmt.Fprintln(bw, lineStr)
		if err != nil {
			return err
		}
	}
	return bw.Flush()
}
