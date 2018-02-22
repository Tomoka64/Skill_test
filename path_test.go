package main

import (
	"runtime"
	"testing"
)

//
func TestImportPkg(t *testing.T) {

	dir := getDir()
	tests := []struct {
		Fname string
		Dir   string
	}{
		{Fname: "fmt", Dir: dir},
		{Fname: "math", Dir: dir},
		{Fname: "strings", Dir: dir},
		{Fname: "bytes", Dir: dir},
	}
	for _, test := range tests {
		p, err := importPkg(test.Fname, test.Dir)
		if err != nil {
			t.Fatalf("expected: %v but got %v", err, nil)
		}
		actual := p.Dir
		expected := runtime.GOROOT() + "/src/" + test.Fname

		if actual != expected {
			t.Fatalf("expected: %v but got %v", expected, actual)
		}
	}
}
