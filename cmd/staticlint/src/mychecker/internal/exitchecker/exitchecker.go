package exitchecker

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var ExitCheckAnalyzer = &analysis.Analyzer{
	Name: "exitcheck",
	Doc:  "check for using os.Exit func in main",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		if file.Name.Name != "main" {
			return nil, nil
		}

		ast.Inspect(file, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.CallExpr:
				selexpr, ok := x.Fun.(*ast.SelectorExpr)
				if !ok {
					return true
				}
				ident, ok := selexpr.X.(*ast.Ident)
				if !ok || ident.Name != "os" {
					return true
				}
				if selexpr.Sel.Name == "Exit" {
					pass.Reportf(selexpr.Pos(), "call os.Exit in main package")
				}
			}
			return true
		})
	}
	return nil, nil
}
