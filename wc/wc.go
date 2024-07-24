package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	chars, words, lines := 0, 0, 0
	reader := strings.NewReader("")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines++
		chars += len(scanner.Text()) + 1
		reader.Reset(scanner.Text())
		wordscanner := bufio.NewScanner(reader)
		wordscanner.Split(bufio.ScanWords)
		for wordscanner.Scan() {
			words++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	} else {
		fmt.Println(lines, words, chars)
	}
}
