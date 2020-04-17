package main

import (
	"os"
)

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	errorCheck(app.Run(os.Args))
}
