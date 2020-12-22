package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "bad invoke: supply one arg to be prepended")
		os.Exit(1)
	}

	os.Stdout.Write([]byte(os.Args[1]))

	if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
