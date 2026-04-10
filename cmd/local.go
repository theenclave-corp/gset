package cmd

import (
	"errors"
	"fmt"

	"github.com/theenclave-corp/gset/internal/config"
	"github.com/theenclave-corp/gset/internal/prompt"
	"github.com/spf13/cobra"
)

var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Configure .git/config for the current repository (name and email only)",
	RunE:  runLocal,
}

var (
	localName            string
	localEmail           string
	localEditor          string // accepted but unused; editor is always global
	localYes             bool
	localDryRun          bool
	localNoBestPractices bool // accepted but unused; local never applies best practices
)

func init() {
	localCmd.Flags().StringVar(&localName, "name", "", "git user.name")
	localCmd.Flags().StringVar(&localEmail, "email", "", "git user.email")
	localCmd.Flags().StringVar(&localEditor, "editor", "", "ignored for local scope; use 'gset global' to set editor")
	localCmd.Flags().BoolVarP(&localYes, "yes", "y", false, "skip confirmation")
	localCmd.Flags().BoolVar(&localDryRun, "dry-run", false, "show what would be applied, make no changes")
	localCmd.Flags().BoolVar(&localNoBestPractices, "no-best-practices", false, "ignored for local scope; best practices are always global-only")
	rootCmd.AddCommand(localCmd)
}

func runLocal(_ *cobra.Command, _ []string) error {
	id := &prompt.Identity{
		Name:  localName,
		Email: localEmail,
		// Editor intentionally omitted — never written to local config
	}

	// withEditor=false: no editor prompt, editor not written locally
	if err := prompt.AskIdentity(id, false); err != nil {
		return err
	}

	scope := config.LocalScope()
	settings := buildSettings(id, false)

	printSummary(".git/config", settings)

	if localDryRun {
		fmt.Println("(dry run — no changes made)")
		return nil
	}

	if !localYes {
		ok, err := prompt.Confirm("Apply?")
		if err != nil {
			if errors.Is(err, prompt.ErrAborted) {
				fmt.Println("Aborted.")
				return nil
			}
			return err
		}
		if !ok {
			fmt.Println("Aborted.")
			return nil
		}
	}

	return applySettings(scope, settings)
}
