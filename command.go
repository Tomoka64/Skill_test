package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Tomoka64/go-skilltest"
	"github.com/fatih/color"
)

//CommandLine is for the command-tool usage. gets File and pattern from the user and holds data to return accordingly.
type CommandLine struct {
	Path    string
	File    string
	Pattern string
	datas   []model.Result
}

func newCommandLine(items ...string) (Driver, error) {
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

//Run will writes data into terminal and implements SaveToFile so that
// the stored data can be used when the usage 'history' has been selected.
func (c *CommandLine) Run() error {
	if err := c.Extract(); err != nil {
		return err
	}
	fmt.Println("Directory\tKeyword\t\tLine\tDetail")
	for _, v := range c.datas {
		color.Yellow("%v\t%v\t%d\t%v\n", v.Filename, v.Keyword, v.Line, v.Detail)
		SaveToFile(&v)
	}
	return nil
}

//CLWord extracts the result according to the given filaname and pattern and puts it into datas.
func (c *CommandLine) CLWord(fname string) error {
	datas, err := extractWord(fname, c.Pattern, c.datas)
	if err != nil {
		return err
	}
	c.datas = datas
	return nil
}

//Extract finds the right path to the expected file and implements CLWord
func (c *CommandLine) Extract() error {
	p, err := importPkg(c.File, c.Path)
	if err != nil {
		return err
	}

	for _, f := range p.GoFiles {
		fname := filepath.Join(p.Dir, f)
		err = c.CLWord(fname)
		if err != nil {
			return err
		}
	}
	return nil
}
