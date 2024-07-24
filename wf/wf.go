package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/scanner"
)

func main() {
	minLen := 3
	frequencies := map[string]int{}
	process := func(s scanner.Scanner) {
		if len(s.TokenText()) >= minLen {
			frequencies[strings.ToLower(s.TokenText())]++
		}
	}
	if len(os.Args) < 2 { // read from stdin
		ScanTokens(os.Stdin, "", process)
		for k, v := range frequencies {
			fmt.Printf("%d %s\n", v, k)
		}
		return
	}
	for _, fileName := range os.Args[1:] {
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		ScanTokens(f, fileName, process)
		fmt.Printf("%s:\n", fileName)
		for k, v := range frequencies {
			fmt.Printf("%d %s\n", v, k)
		}
		f.Close()
		clear(frequencies)
	}
}

type processToken func(scanner.Scanner)

func ScanTokens(src io.Reader, fileName string, process processToken) {
	var s scanner.Scanner
	s.Init(src)
	s.Filename = fileName
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		if tok == scanner.Ident {
			process(s)
		}
	}
}
