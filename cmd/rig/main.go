package main

import (
	"fmt"
	"os"
)

const usage = `rig - personal environment manager

Usage:
  rig sync              deploy dotfiles from repo to ~/
  rig push [message]    copy dotfile changes from ~/ back to repo and push
  rig upgrade           download latest rig binary from GitHub releases
`

func main() {
	if len(os.Args) < 2 {
		fmt.Print(usage)
		os.Exit(0)
	}

	var err error
	switch os.Args[1] {
	case "sync":
		err = runSync()
	case "push":
		msg := "chore: update dotfiles"
		if len(os.Args) >= 3 {
			msg = os.Args[2]
		}
		err = runPush(msg)
	case "upgrade":
		err = runUpgrade()
	case "help", "--help", "-h":
		fmt.Print(usage)
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n%s", os.Args[1], usage)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
