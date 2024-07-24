package tokens

import (
	"io"
	"text/scanner"
)

type ProcessToken func(scanner.Scanner)

func Scan(src io.Reader, fileName string, process ProcessToken) {
	var s scanner.Scanner
	s.Init(src)
	s.Filename = fileName
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		if tok == scanner.Ident {
			process(s)
		}
	}
}
