package cmd

import (
	"errors"
	"fmt"

	"github.com/theenclave-corp/gset/internal/config"
	"github.com/theenclave-corp/gset/internal/prompt"
	"github.com/spf13/cobra"
)

var globalCmd = &cobra.Command{
	Use:   "global",
	Short: "Configure ~/.gitconfig with best-practice defaults",
	RunE:  runGlobal,
}

var (
	globalName            string
	globalEmail           string
	globalEditor          string
	globalYes             bool
	globalDryRun          bool
	globalNoBestPractices bool
)

func init() {
	globalCmd.Flags().StringVar(&globalName, "name", "", "git user.name")
	globalCmd.Flags().StringVar(&globalEmail, "email", "", "git user.email")
	globalCmd.Flags().StringVar(&globalEditor, "editor", "", "editor: vscode, vim, nano, emacs")
	globalCmd.Flags().BoolVarP(&globalYes, "yes", "y", false, "skip confirmation")
	globalCmd.Flags().BoolVar(&globalDryRun, "dry-run", false, "show what would be applied, make no changes")
	globalCmd.Flags().BoolVar(&globalNoBestPractices, "no-best-practices", false, "only set identity, skip best-practice settings")
	rootCmd.AddCommand(globalCmd)
}

func runGlobal(_ *cobra.Command, _ []string) error {
	id := &prompt.Identity{
		Name:   globalName,
		Email:  globalEmail,
		Editor: globalEditor,
	}

	if err := prompt.AskIdentity(id, true); err != nil {
		return err
	}

	scope := config.GlobalScope()
	settings := buildSettings(id, !globalNoBestPractices)

	printSummary("~/.gitconfig", settings)

	if globalDryRun {
		fmt.Println("(dry run — no changes made)")
		return nil
	}

	if !globalYes {
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
