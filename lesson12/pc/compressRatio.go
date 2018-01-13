package main

import (
	"fmt"
	"compress/zlib"
	"bytes"
	"io"
	"strings"
	"os"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

type GzipCounter struct {
	c int
	Z *zlib.Writer
}
func (g *GzipCounter) Write(p []byte) (int, error) {
	var buf bytes.Buffer
	g.Z = zlib.NewWriter(&buf)
	io.Copy(g.Z, strings.NewReader(string(p)))
	g.Z.Flush()
	var c ByteCounter
	io.Copy(&c, &buf)
	g.c += int(c)
	return len(p), nil
}


func main() {
	var bcgz ByteCounter
	var gzc GzipCounter
	//tee := io.TeeReader(os.Stdin, &gzc)

	io.Copy(&bcgz, io.TeeReader(os.Stdin, &gzc))
	fmt.Printf("gzip: %f\n", 1 - float64(gzc.c)/float64(bcgz))
}