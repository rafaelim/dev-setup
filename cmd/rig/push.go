package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func runPush(msg string) error {
	rigDir, err := getRigDir()
	if err != nil {
		return err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	envDir := filepath.Join(rigDir, "env")
	fmt.Println("Collecting dotfile changes...")

	err = filepath.Walk(envDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		rel, _ := filepath.Rel(envDir, path)
		src := envRelToHome(rel, home)

		if _, err := os.Stat(src); os.IsNotExist(err) {
			fmt.Printf("  skipping (not found): %s\n", src)
			return nil
		}
		fmt.Printf("  <- %s\n", src)
		return copyFile(src, path, info.Mode())
	})
	if err != nil {
		return err
	}

	out, _ := exec.Command("git", "-C", rigDir, "status", "--porcelain", "env/").Output()
	if len(strings.TrimSpace(string(out))) == 0 {
		fmt.Println("Nothing to push.")
		return nil
	}

	fmt.Println("Pushing to GitHub...")
	if err := gitRun(rigDir, "add", "env/"); err != nil {
		return err
	}
	if err := gitRun(rigDir, "commit", "-m", msg); err != nil {
		return err
	}
	return gitRun(rigDir, "push")
}
