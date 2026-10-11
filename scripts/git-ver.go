package main

import (
	"Dbot/console"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
)

type VersionConfig struct {
	Information VersionInformation    `yaml:"information"`
	Git         VersionGitInformation `yaml:"git"`
}

type VersionInformation struct {
	AppName string `yaml:"app_name"`
	Version string `yaml:"version"`
}

type VersionGitInformation struct {
	Hash  string `yaml:"hash"`
	Title string `yaml:"title"`
}

func main() {
	cmd := exec.Command("git", "log", "-1", `pretty=format:"%h|%s"`)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		console.Error(fmt.Sprintf("Failed to run git command... %v", err))
		return
	}
	output := strings.TrimSpace(out.String())
	parts := strings.SplitN(output, "|", 2)
	if len(parts) < 2 || parts[0] == "" || parts[1] == "" {
		console.Error("Couldn't find commit or hash?")
		return
	}
	hash, title := parts[0], parts[1]
	cwd, err := os.Getwd()
	if err != nil {
		console.Error(fmt.Sprintf("Failed to get cwd... %v", err))
		return
	}
	path := filepath.Join(cwd, "VERSION.yml")
	raw, err := os.ReadFile(path)
	if err != nil {
		console.Error(fmt.Sprintf("Failed to read file... %v", err))
		return
	}
	var parsed VersionConfig
	if err := yaml.Unmarshal(raw, &parsed); err != nil {
		console.Error(fmt.Sprintf("Failed to parse... %v", err))
		return
	}

	if parsed.Git.Hash == hash {
		console.Info("Up-to-date")
		return
	}

	parsed.Git.Hash = hash
	parsed.Git.Title = title

	updRaw, err := yaml.Marshal(&parsed)
	if err != nil {
		console.Error(fmt.Sprintf("Failed to marshal back into yml... %v", err))
		return
	}
	if err := os.WriteFile(path, updRaw, 0o644); err != nil { // 0 / 6 (owner) / 4 (group) / 4 (others);
		console.Error(fmt.Sprintf("Failed to write... %v", err))
		return
	}

	console.Info("Successfully updated version")
}

/*
	with octal permission values, it's quite simple
	0o (octal marker for golang) | 6 (owner) | 4 (group) | 4 (others)
	group: completely useless on windows
	other: same as group, useless on windows
	owner: the permission digit for the owner of the file

	r = read; w = write; x = execute
	7 = rwx
	6 = rw-
	5 = r-x
	4 = r--
	0 = ---
*/
