package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const releaseURL = "https://github.com/rafaelim/dev-setup/releases/latest/download/rig"

func runUpgrade() error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	exe, err = filepath.EvalSymlinks(exe)
	if err != nil {
		return err
	}

	fmt.Println("Downloading latest rig from GitHub releases...")
	resp, err := http.Get(releaseURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: %s", resp.Status)
	}

	tmp := exe + ".tmp"
	f, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	if _, err := io.Copy(f, resp.Body); err != nil {
		f.Close()
		os.Remove(tmp)
		return err
	}
	f.Close()

	if err := os.Rename(tmp, exe); err != nil {
		os.Remove(tmp)
		return err
	}

	fmt.Println("rig upgraded successfully.")
	return nil
}
