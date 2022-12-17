package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "test/root.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Comments:")
	for _, c := range node.Comments {
		fmt.Print(c.Text())
	}

	fmt.Println("Functions:")
	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}
		fmt.Println(fn.Name.Name)

		ast.Inspect(f, func(node ast.Node) bool {
			switch n := node.(type) {
			case *ast.CallExpr:
				fmt.Println(n.Fun) // prints every func call expression

				if fun, ok := n.Fun.(*ast.SelectorExpr); ok {
					fmt.Println(fun.Sel.Name)
					fmt.Println(fun.Sel.String())
				}

				for _, arg := range node.(*ast.CallExpr).Args {
					fmt.Println(arg)
				}
			}
			return true
		})
	}
}
