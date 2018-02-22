package main

import (
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	tests := []struct {
		args []string
		ans  Driver
	}{
		{
			args: []string{"localhost"},
			ans:  &Server{},
		},
		{
			args: []string{"history"},
			ans:  &History{},
		},
		{
			args: []string{"fmt", "TODO"},
			ans:  &CommandLine{},
		},
		{
			args: []string{},
			ans:  &Helper{},
		},
		{
			args: []string{""},
			ans:  &Helper{},
		},
		{
			args: []string{"1", "2", "3"},
			ans:  &Helper{},
		},
	}

	for _, test := range tests {
		n, err := New(test.args)
		if err != nil {
			t.Fatalf("expected %v but got nil", err)
		}
		if reflect.TypeOf(n) != reflect.TypeOf(test.ans) {
			t.Fatalf("expected %v but got %v", n, test.ans)
		}
	}
}
