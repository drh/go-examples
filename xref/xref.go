package main

import (
	"fmt"
	"go-examples/internal/sortmapkeys"
	"go-examples/internal/tokens"
	"go/token"
	"io"
	"os"
	"text/scanner"
)

func main() {
	// identifiers[id] is a map[filename] of the files in which id appears.
	// The values in this map are int slices that hold the line numbers in filename
	// in which id appears. Since the files are read sequentially, the line numbers
	// in the slices are sorted in ascending order.
	identifiers := map[string]map[string][]int{}
	for _, fileName := range os.Args[1:] {
		f, err := os.Open(fileName)
		if err == nil {
			xref(f, fileName, identifiers)
			f.Close()
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	if len(os.Args) == 1 { // read from stdin
		xref(os.Stdin, "", identifiers)
	}

	// identifiers has been populated. The output has the form
	// id1 filename1 linenumbers...
	//     filename2 linenumbers...
	//     ...
	// id2 filename2 linenumbers...
	//     filename3 linenumbers...
	// ...
	// The ids and files names are sorted in ascending order.
	//
	// If there are no program arguments, the output has the form
	// id1 linenumbers...
	// id2 linenumbers...
	// ...
	//
	ids := sortmapkeys.SortByKey(identifiers)
	for _, id := range ids {
		fmt.Printf("%s", id)
		print(identifiers[id])
	}
}

func print(fileNames map[string][]int) {
	if lineNumbers, ok := fileNames[""]; ok { // input read from stdin
		printLineNumbers(lineNumbers, "\t")
	} else {
		names := sortmapkeys.SortByKey(fileNames)
		for _, name := range names {
			fmt.Printf("\t%s", name)
			printLineNumbers(fileNames[name], " ")
		}
	}
}

func printLineNumbers(lineNumbers []int, sep string) {
	for _, ln := range lineNumbers {
		fmt.Printf("%s%d", sep, ln)
		sep = " "
	}
	fmt.Println()
}

func xref(src io.Reader, fileName string, identifiers map[string]map[string][]int) {
	tokens.ScanIdentifiers(src, fileName, func(s scanner.Scanner) {
		id := s.TokenText()
		if token.IsKeyword(id) {
			return
		}
		if m, ok := identifiers[id]; ok { // Another use of id
			if ln, ok := m[fileName]; ok { // ... in this file
				if s.Line != ln[len(ln)-1] {
					m[fileName] = append(ln, s.Line)
				}
			} else { // First use of id in this file
				identifiers[id][fileName] = append(make([]int, 0, 20), s.Line)
			}
		} else { // First use of id anywhere
			m = make(map[string][]int)
			m[fileName] = append(make([]int, 0, 20), s.Line)
			identifiers[id] = m
		}
	})
}
