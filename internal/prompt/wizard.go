package prompt

import (
	"github.com/charmbracelet/huh"
)

// Identity holds the user-supplied identity fields.
type Identity struct {
	Name   string
	Email  string
	Editor string
}

// AskIdentity interactively prompts for identity fields.
// Pre-populated fields are skipped. withEditor controls whether the
// editor prompt is included (true for global, false for local).
func AskIdentity(id *Identity, withEditor bool) error {
	var fields []huh.Field

	if id.Name == "" {
		fields = append(fields, huh.NewInput().
			Title("Name").
			Placeholder("Jane Doe").
			Value(&id.Name))
	}

	if id.Email == "" {
		fields = append(fields, huh.NewInput().
			Title("Email").
			Placeholder("jane@example.com").
			Value(&id.Email))
	}

	if withEditor && id.Editor == "" {
		fields = append(fields, huh.NewSelect[string]().
			Title("Editor").
			Options(
				huh.NewOption("VS Code", "vscode"),
				huh.NewOption("Vim", "vim"),
				huh.NewOption("Nano", "nano"),
				huh.NewOption("Emacs", "emacs"),
			).
			Value(&id.Editor))
	}

	if len(fields) == 0 {
		return nil
	}

	return huh.NewForm(huh.NewGroup(fields...)).Run()
}

// Confirm shows a yes/no prompt with the given message.
// Returns true if the user confirms.
func Confirm(msg string) (bool, error) {
	var ok bool
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(msg).
				Value(&ok),
		),
	).Run()
	return ok, err
}
