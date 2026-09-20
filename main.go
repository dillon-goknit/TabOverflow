package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "taboverflow:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) < 1 {
		return errors.New("no command given")
	}

	switch args[0] {
	case "add":
		if len(args) < 2 {
			return errors.New("add requires a URL")
		}
		return cmdAdd(args[1])
	case "list":
		cmdList()
	case "pick":
		cmdPick()
	case "done":
		if len(args) < 2 {
			return errors.New("done requires a URL")
		}
		cmdDone(args[1])
	case "rm":
		if len(args) < 2 {
			return errors.New("rm requires a URL")
		}
		cmdRm(args[1])
	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}

	return nil
}
