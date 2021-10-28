package main

import (
	"os"
)

func main() {
	startApp(os.Stdin, os.Stdout)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
