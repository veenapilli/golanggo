package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Expected file name to print as the only argument")
		os.Exit(1)
	}
	t := args[1]
	f, err := os.Open(t)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
