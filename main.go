package main

import (
	"log"
	"os"
)

type another interface {
	Run() error
}

//New determines the usage that user wants to use.
//there are 3 usages; command-tool purposed usage(to simply search the selected word in selected package),
//'localhost' usage (to make server-mode possible), and  'history' usage(to show all the searched history).
//if the user's request does not fit any of the usage above, it goes to newHelper(to show users
// the usage of this command tool)
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
