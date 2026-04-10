package config

import (
	"fmt"
	"os/exec"
	"strings"
)

// Set writes a single git config key=value using the given scope args.
// scope is e.g. []string{"--global"}, []string{"--local"}, or []string{"--file", path}.
func Set(scope []string, key, value string) error {
	args := append([]string{"config"}, scope...)
	args = append(args, key, value)
	out, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("git config %s %s: %w\n%s", key, value, err, out)
	}
	return nil
}

// Get reads a single git config key using the given scope args.
// Returns an error if the key does not exist.
func Get(scope []string, key string) (string, error) {
	args := append([]string{"config"}, scope...)
	args = append(args, key)
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return "", fmt.Errorf("key %q not found: %w", key, err)
	}
	return strings.TrimRight(string(out), "\n"), nil
}

// List returns all key=value pairs in the given scope as a map.
// Returns an empty map (not an error) when the config file is empty or absent.
func List(scope []string) (map[string]string, error) {
	args := append([]string{"config"}, scope...)
	args = append(args, "--list")
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return map[string]string{}, nil
	}
	result := map[string]string{}
	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			result[parts[0]] = parts[1]
		}
	}
	return result, nil
}

// GlobalScope returns the scope args for ~/.gitconfig.
func GlobalScope() []string { return []string{"--global"} }

// LocalScope returns the scope args for the repo's .git/config.
func LocalScope() []string { return []string{"--local"} }
