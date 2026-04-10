package cmd

import (
	"fmt"
	"sort"

	"github.com/theenclave-corp/gset/internal/config"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current effective git configuration",
	RunE:  runStatus,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func runStatus(_ *cobra.Command, _ []string) error {
	global, err := config.List(config.GlobalScope())
	if err != nil {
		return fmt.Errorf("reading global config: %w", err)
	}

	// Local config is optional — not inside a git repo is fine
	local, _ := config.List(config.LocalScope())

	// Merge: local overrides global
	merged := map[string]string{}
	for k, v := range global {
		merged[k] = v
	}
	for k, v := range local {
		merged[k] = v
	}

	if len(merged) == 0 {
		fmt.Println("No git configuration found.")
		return nil
	}

	keys := make([]string, 0, len(merged))
	for k := range merged {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("Effective git configuration:")
	fmt.Println()
	for _, k := range keys {
		origin := ""
		if _, inLocal := local[k]; inLocal {
			origin = " (local)"
		}
		fmt.Printf("  %-35s = %s%s\n", k, merged[k], origin)
	}
	return nil
}
