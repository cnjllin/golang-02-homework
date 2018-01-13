package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode/utf8"
)

func isSpace(r rune)bool {
	if r == ',' || r == ' ' {
		return true
	}
	return false
}

func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !isSpace(r) {
			break
		}
	}
	// Scan until space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if isSpace(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	buf := strings.NewReader(string(p))
	s := bufio.NewScanner(buf)
	s.Split(ScanWords)
	for s.Scan() {
		fmt.Println(s.Text())
		*c ++
	}
	return len(p), nil
}

func main() {
	var bc WordCounter
	fmt.Fprintf(&bc, "   fdsafdsa, dsaf df  fdsaf,fdsafdsaffdsaf")
	fmt.Println(bc)
	//fmt.Println(bufio.ScanWords([]byte("hello worlsssd"), false))
}