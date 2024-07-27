package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/scanner"

	"github.com/drh/go-examples/internal/sortmapkeys"
	"github.com/drh/go-examples/internal/tokens"
)

const minLength = 3

func main() {
	for _, fileName := range os.Args[1:] {
		f, err := os.Open(fileName)
		if err == nil {
			wf(f, fileName)
			f.Close()
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	if len(os.Args) == 1 { // read from stdin
		wf(os.Stdin, "")
	}
}

func wf(src io.Reader, fileName string) {
	frequencies := map[string]int{}
	tokens.ScanWords(src, fileName, func(s scanner.Scanner) {
		if len(s.TokenText()) >= minLength {
			frequencies[strings.ToLower(s.TokenText())]++
		}
	})
	if fileName != "" {
		fmt.Printf("%s:\n", fileName)
	}
	// Sort on values
	keys := sortmapkeys.SortByValue(frequencies)
	for _, k := range keys {
		fmt.Printf("%d %s\n", frequencies[k], k)
	}
}
