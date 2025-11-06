package tools

import (
	"fmt"
	"os"
	"path/filepath"
)

func listEmacsServers() ([]string, error) {
	// Get XDG_RUNTIME_DIR from environment
	runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
	if runtimeDir == "" {
		return nil, fmt.Errorf("XDG_RUNTIME_DIR not set")
	}

	emacsDir := filepath.Join(runtimeDir, "emacs")

	entries, err := os.ReadDir(emacsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", emacsDir, err)
	}

	files := []string{}
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, filepath.Join(emacsDir, entry.Name()))
		}
	}

	return files, nil
}
