package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Tomoka64/final/model"
	"github.com/fatih/color"
)

type CommandLine struct {
	Path    string
	File    string
	Pattern string
	datas   []model.Result
}

func newCommandLine(items ...string) (another, error) {
	dir, err := os.Getwd()
	if err != nil {
		return &CommandLine{}, err
	}
	return &CommandLine{
		Path:    dir,
		File:    items[0],
		Pattern: items[1],
	}, nil
}

func (r *CommandLine) Run() error {
	if err := r.Extract(); err != nil {
		return err
	}
	fmt.Println("Directory\tKeyword\t\tLine\tDetail")
	for _, v := range r.datas {
		color.Yellow("%v\t%v\t%d\t%v\n", v.Filename, v.Keyword, v.Line, v.Detail)
		SaveToFile(&v)
	}
	return nil
}

func (r *CommandLine) CLWord(fname string) error {
	datas, err := extractWord(
		fname, r.Pattern, r.datas)
	if err != nil {
		return err
	}
	r.datas = datas
	return nil
}

func (r *CommandLine) Extract() error {
	p, err := importPkg(r.File, r.Path)
	if err != nil {
		return err
	}

	for _, f := range p.GoFiles {
		fname := filepath.Join(p.Dir, f)
		err = r.CLWord(fname)
		if err != nil {
			return err
		}
	}
	return nil
}
