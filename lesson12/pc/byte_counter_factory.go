package main

import (
	"fmt"
	"io"
	//"bytes"
	"os"
	"github.com/siddontang/go/num"
	"strings"
)

//!+bytecounter

type ByteCnt struct {
	Cnt int64
	w   io.Writer
}

func (c *ByteCnt) Write(p []byte) (int, error) {
	//io.Copy(c.w, bytes.NewReader(p))
	n, err := c.w.Write(p)
	c.Cnt += int64(len(p)) // convert int to ByteCounter
	return n, err
}

//!-bytecounter
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	nw := ByteCnt{0, w}

	return &nw, &nw.Cnt
}

type LimitReader struct {
	cnt int
	r io.Reader
}

func (l * LimitReader) Read(p []byte) (n int, err error) {
	cnt := num.MinInt(len(p), l.cnt)
	fmt.Println(cnt)
	return l.r.Read(p[0:cnt])
}

func NewLimitReader(r io.Reader, max int)(io.Reader) {
	nr := LimitReader{max, r}
	return &nr
}

func main() {
	//!+main
	nc, count_p := CountingWriter(os.Stderr)
	r := strings.NewReader("afdsafdsafdsafdsa")
	nr := NewLimitReader(r, 2)
	io.Copy(os.Stdout, nr)

	var name = "Dolly"
	fmt.Fprintf(nc, "hello, %s", name)
	//fmt.Println(nc.Cnt) // "12", = len("hello, Dolly")
	fmt.Println(*count_p)
	//!-main
}
