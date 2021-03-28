package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "sample/sample.go", nil, 0)
	if err != nil {
		panic(err)
	}

	ast.Print(fset, astFile)

}
