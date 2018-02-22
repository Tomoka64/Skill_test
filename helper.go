package main

import (
	"fmt"
	"os"
)

type Helper struct {
	message string
}

func newHelper(items ...string) (another, error) {
	return &Helper{
		fmt.Sprintf("usage1: %s <directory> <keyword> (e.g. fmt TODO)\nusage2: localhost (to connect with localhost)\nusage3: history (to see your search history)", items[0]),
	}, nil
}

func (r *Helper) Run() error {
	fmt.Fprintln(os.Stdout, r.message)
	os.Exit(1)
	return nil
}
