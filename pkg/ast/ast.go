package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)


func Init() {
	src := `
package main

func main() {
	println("Hello World!")
}
`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	ast.Print(fset, f)
}	

func main() {
	fmt.Println(parse(`(a == "123") && (b == "456")`))	
}


func parse(expr string) (bool, error) {
	exprAst, err := parser.ParseExpr(expr)
	if err != nil {
		return false, nil		
	}

	return judge(exprAst), nil
}


// 二元表达式 a > b

func judge(bop ast.Node) bool {
	// 断言成二元表达式 

	expr := bop.(*ast.BinaryExpr)
	x := expr.X.(*ast.BinaryExpr)
	key := x.Y.(*ast.BasicLit)  // key 值

	y := expr.Y.(*ast.BinaryExpr)
	value := y.Y.(*ast.BasicLit)  // value 值
	

	return key.Value == value.Value
}