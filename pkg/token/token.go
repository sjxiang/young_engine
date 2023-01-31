package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	var src = []byte(`price > 500 && (isNewUser || userLevel > 5)`)  // println("Hello World!")

	var fset = token.NewFileSet()
	var file = fset.AddFile("hello.go", fset.Base(), len(src))

	var s scanner.Scanner
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()  // 索引，token 标识，字面量
		if tok == token.EOF {
			break
		}

		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}