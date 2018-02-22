package main

import (
	"log"
	"os"
)

type another interface {
	Run() error
}

func New(args []string) (another, error) {
	var f func(item ...string) (another, error)

	switch len(args) {
	case 1:
		switch args[0] {
		case "localhost":
			f = newServer
		case "history":
			f = newHistory
		default:
			f = newHelper
			args = os.Args
		}
	case 2:
		f = newCommandLine
	default:
		f = newHelper
		args = os.Args
	}
	return f(args...)
}

func main() {
	t, err := New(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	err = t.Run()
	if err != nil {
		log.Fatal(err)
	}
}
