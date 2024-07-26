package tokens

import (
	"io"
	"text/scanner"
	"unicode"
)

type ProcessToken func(scanner.Scanner)

func Scan(s scanner.Scanner, src io.Reader, process ProcessToken) {
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		if tok == scanner.Ident {
			process(s)
		}
	}
}

func ScanWords(src io.Reader, fileName string, process ProcessToken) {
	var s scanner.Scanner
	s.Init(src)
	s.Filename = fileName
	s.Mode = scanner.ScanIdents | scanner.ScanComments
	s.IsIdentRune = func(ch rune, i int) bool { return unicode.IsLetter(ch) }
	Scan(s, src, process)
}

func ScanIdentifiers(src io.Reader, fileName string, process ProcessToken) {
	var s scanner.Scanner
	s.Init(src)
	s.Filename = fileName
	s.Mode = scanner.GoTokens
	Scan(s, src, process)
}
