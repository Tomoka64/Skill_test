package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCommandLine(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("%v", err)
	}
	items := []struct {
		item   []string
		result CommandLine
		err    error
	}{
		{item: []string{"fmt", "TODO"},
			result: CommandLine{
				Path:    dir,
				File:    "fmt",
				Pattern: "TODO",
			},
			err: nil},
		{item: []string{"math", "TODO"},
			result: CommandLine{
				Path:    dir,
				File:    "math",
				Pattern: "TODO",
			},
			err: nil},
		{item: []string{"strings", "TODO"},
			result: CommandLine{
				Path:    dir,
				File:    "strings",
				Pattern: "TODO",
			},
			err: nil},
		{item: []string{"strings", "NOTE"},
			result: CommandLine{
				Path:    dir,
				File:    "strings",
				Pattern: "NOTE",
			},
			err: nil},
		{item: []string{"runtime", ""},
			result: CommandLine{
				Path:    dir,
				File:    "runtime",
				Pattern: "",
			},
			err: err},
		{item: []string{"", ""},
			result: CommandLine{
				Path:    dir,
				File:    "",
				Pattern: "",
			},
			err: err},
	}
	for _, test := range items {
		result, err := newCommandLine(test.item[0], test.item[1])
		assert.IsType(t, test.err, err)
		assert.Equal(t, &test.result, result)
	}

}

func getDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	return dir
}

func TestRun(t *testing.T) {
	var err error
	items := []struct {
		result CommandLine
		err    error
	}{
		{result: CommandLine{
			Path:    getDir(),
			File:    "fmt",
			Pattern: "TODO",
		},
			err: nil},
		{result: CommandLine{
			Path:    getDir(),
			File:    "math",
			Pattern: "TODO",
		},
			err: nil},
		{result: CommandLine{
			Path:    getDir(),
			File:    "strings",
			Pattern: "TODO",
		},
			err: nil},
		{result: CommandLine{
			Path:    getDir(),
			File:    "",
			Pattern: "",
		},
			err: err},
	}
	for _, test := range items {
		err := test.result.Run()
		if err == nil {
			assert.IsType(t, test.err, err)

		}

	}

}

func TestCLWord(t *testing.T) {
	var err error
	tests := []struct {
		fname string
		err   error
	}{
		{fname: "fmt", err: nil},
		{fname: "math", err: nil},
		{fname: "strings", err: nil},
		{fname: "bytes", err: nil},
		{fname: "runtime", err: nil},
		{fname: "", err: err},
	}
	for _, test := range tests {
		var r CommandLine
		err := r.CLWord(test.fname)
		if err == nil {
			assert.IsType(t, test.err, err)
		}

	}
}
