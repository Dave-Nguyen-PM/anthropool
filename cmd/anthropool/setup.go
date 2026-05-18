package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Dave-Nguyen-PM/anthropool/internal/atomicfile"
	"github.com/Dave-Nguyen-PM/anthropool/internal/paths"
)

// switchSlashCommand is the markdown body installed at
// ~/.claude/commands/switch.md. Embedding keeps the binary self-contained
// so `anthropool setup` works on a freshly downloaded binary with no source tree.
//
//go:embed slashcmd_switch.md
var switchSlashCommand []byte

//go:embed slashcmd_anthropool_add.md
var anthropoolAddSlashCommand []byte

//go:embed slashcmd_anthropool_list.md
var anthropoolListSlashCommand []byte

//go:embed slashcmd_anthropool_status.md
var anthropoolStatusSlashCommand []byte

//go:embed slashcmd_anthropool_support.md
var anthropoolSupportSlashCommand []byte

//go:embed slashcmd_anthropool_switch.md
var anthropoolSwitchSlashCommand []byte

//go:embed slashcmd_anthropool_config.md
var anthropoolConfigSlashCommand []byte

//go:embed slashcmd_anthropool_remove.md
var anthropoolRemoveSlashCommand []byte

//go:embed slashcmd_anthropool_usage_refresh.md
var anthropoolUsageRefreshSlashCommand []byte

func installSlashCommand() error {
	dir := filepath.Join(paths.ClaudeDir(), "commands")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return fmt.Errorf("setup: mkdir %s: %w", dir, err)
	}
	if err := atomicfile.Write(filepath.Join(dir, "switch.md"), switchSlashCommand, 0o600); err != nil {
		return err
	}

	anthropoolDir := filepath.Join(dir, "anthropool")
	if err := os.MkdirAll(anthropoolDir, 0o700); err != nil {
		return fmt.Errorf("setup: mkdir %s: %w", anthropoolDir, err)
	}
	commands := map[string][]byte{
		"add.md":           anthropoolAddSlashCommand,
		"config.md":        anthropoolConfigSlashCommand,
		"list.md":          anthropoolListSlashCommand,
		"status.md":        anthropoolStatusSlashCommand,
		"support.md":       anthropoolSupportSlashCommand,
		"switch.md":        anthropoolSwitchSlashCommand,
		"remove.md":        anthropoolRemoveSlashCommand,
		"usage-refresh.md": anthropoolUsageRefreshSlashCommand,
	}
	for name, body := range commands {
		if err := atomicfile.Write(filepath.Join(anthropoolDir, name), body, 0o600); err != nil {
			return err
		}
	}
	return nil
}
