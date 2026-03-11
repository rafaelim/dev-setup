package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func runSync() error {
	rigDir, err := getRigDir()
	if err != nil {
		return err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	envDir := filepath.Join(rigDir, "env")
	fmt.Println("Syncing dotfiles...")

	return filepath.Walk(envDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		rel, _ := filepath.Rel(envDir, path)
		dest := envRelToHome(rel, home)
		fmt.Printf("  -> %s\n", dest)
		return copyFile(path, dest, info.Mode())
	})
}

// envRelToHome maps a relative path under env/ to its home directory destination.
// env/personal/.gitconfig → ~/.gitconfig
// env/.config/nvim/...   → ~/.config/nvim/...
func envRelToHome(rel, home string) string {
	if strings.HasPrefix(rel, "personal/") {
		return filepath.Join(home, strings.TrimPrefix(rel, "personal/"))
	}
	return filepath.Join(home, rel)
}
