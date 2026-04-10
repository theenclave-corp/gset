package cmd

import (
	"fmt"
	"os"

	"github.com/theenclave-corp/gset/internal/config"
	"github.com/theenclave-corp/gset/internal/prompt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gset",
	Short: "Configure git settings following best practices",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// buildSettings composes the full list of settings to apply.
// When includeBestPractices is false, only identity fields are included.
func buildSettings(id *prompt.Identity, includeBestPractices bool) []config.Setting {
	var settings []config.Setting

	if id.Name != "" {
		settings = append(settings, config.Setting{
			Key: "user.name", Value: id.Name,
			Category: "Identity", Description: "committer name",
		})
	}
	if id.Email != "" {
		settings = append(settings, config.Setting{
			Key: "user.email", Value: id.Email,
			Category: "Identity", Description: "committer email",
		})
	}
	if id.Editor != "" {
		settings = append(settings, config.Setting{
			Key: "core.editor", Value: config.EditorValue(id.Editor),
			Category: "Identity", Description: "default editor",
		})
	}

	if includeBestPractices {
		settings = append(settings, config.BestPracticeSettings()...)
	}
	return settings
}

// printSummary prints a grouped, aligned list of settings to be applied.
func printSummary(target string, settings []config.Setting) {
	fmt.Printf("\nApplying %d settings to %s:\n\n", len(settings), target)
	category := ""
	for _, s := range settings {
		if s.Category != category {
			category = s.Category
			fmt.Printf("  %s\n", category)
		}
		fmt.Printf("    %-30s = %s\n", s.Key, s.Value)
	}
	fmt.Println()
}

// applySettings writes each setting using git config.
func applySettings(scope []string, settings []config.Setting) error {
	for _, s := range settings {
		if err := config.Set(scope, s.Key, s.Value); err != nil {
			return fmt.Errorf("failed to set %s: %w", s.Key, err)
		}
		fmt.Printf("  ✓ %s\n", s.Key)
	}
	fmt.Println("\nDone.")
	return nil
}
