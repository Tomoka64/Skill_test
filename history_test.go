package main

import (
	"io/ioutil"
	"testing"

	"github.com/Tomoka64/go-pkg-seeker/model"
	"github.com/stretchr/testify/assert"
)

func TestFileGetContents(t *testing.T) {
	filename := "config/test.txt"
	ioutil.WriteFile(filename, []byte("hello"), 0600)
	expected := "hello"
	actual := FileGetContents(filename)
	if expected != string(actual) {
		t.Fatalf("expected %v but got %v", expected, actual)
	}
}

func TestNewHistory(t *testing.T) {
	tests := []struct {
		item   []string
		result History
		err    error
	}{
		{item: []string{"history"}, result: History{}, err: nil},
	}
	for _, test := range tests {
		result, err := newHistory("history")
		assert.IsType(t, test.err, err)
		assert.IsType(t, &test.result, result)

	}
}

func TestSaveToFile(t *testing.T) {
	tests := []struct {
		m   model.Result
		err error
	}{
		{m: model.Result{
			Filename: "fmt",
			Keyword:  "TODO",
			Detail:   "good",
			Line:     11,
		}, err: nil},
		{m: model.Result{
			Filename: "a",
			Keyword:  "be",
			Detail:   "good",
			Line:     1,
		}, err: nil},
		{m: model.Result{
			Filename: "fmt",
			Keyword:  "r",
			Detail:   "good",
			Line:     1,
		}, err: nil},
		{m: model.Result{
			Filename: "",
			Keyword:  "",
			Detail:   "good",
			Line:     1,
		}, err: nil},
	}

	for _, test := range tests {
		err := SaveToFile(&test.m)
		assert.IsType(t, test.err, err)

	}

}
