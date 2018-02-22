package main

import (
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"strings"

	"github.com/Tomoka64/final/model"
)

//importPkg imports a path to a directory where the fname is located.
func importPkg(fname, dir string) (*build.Package, error) {
	p, err := build.Import(fname, dir, build.ImportComment)
	if err != nil {
		return &build.Package{}, err
	}
	if p.BinaryOnly {
		return &build.Package{}, nil
	}
	if p.IsCommand() {
		return &build.Package{}, nil
	}
	return p, nil
}

//extractWord gets filename, pattern and looks for puts the found results
//into data.
func extractWord(fname, pattern string, datas []model.Result) ([]model.Result, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
	if err != nil {
		return []model.Result{}, err
	}

	cmap := ast.NewCommentMap(fset, f, f.Comments)
	for n, cgs := range cmap {
		f := fset.File(n.Pos())
		for _, cg := range cgs {
			t := cg.Text()
			if strings.Contains(t, pattern) {
				a := f.Position(cg.Pos()).Line
				datas = append(datas, model.NewResult(fname, pattern, t, a))
			}
		}
	}

	return datas, nil
}
