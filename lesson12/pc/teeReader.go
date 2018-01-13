package main

import (
	"io"
	"strings"
	"os"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	tee := io.TeeReader(r, os.Stderr)

	io.Copy(os.Stdout, tee)
}
