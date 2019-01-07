package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing filename")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Stdout.WriteString("asdfasf")

	io.Copy(os.Stdout, file)
}
