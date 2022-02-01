package main

import (
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)

	if err != nil {
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
