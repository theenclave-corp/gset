package config

import (
	"os/exec"
)

// Set runs `git config <scope...> <key> <value>`.
func Set(scope []string, key, value string) error {
	args := append([]string{"config"}, scope...)
	args = append(args, key, value)
	return exec.Command("git", args...).Run()
}
