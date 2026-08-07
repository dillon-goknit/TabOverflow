package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: taboverflow: <command>") // list the commands
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("add requires a URL")
			return
		}
		cmdAdd(os.Args[2])
	case "list":
		cmdList()
	case "pick":
		cmdPick()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("done requires a URL")
			return
		}
		cmdDone(os.Args[2])
	case "rm":
		if len(os.Args) < 3 {
			fmt.Println("rm requires a URL")
			return
		}
		cmdRm(os.Args[2])
	default:
		fmt.Printf("unknown command: %s\n", command)

	}
}
