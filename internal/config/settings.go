package config

// Setting represents a single git config key-value pair with metadata.
type Setting struct {
	Key         string
	Value       string
	Category    string
	Description string
}

// BestPracticeSettings returns the curated list of 18 git settings to apply globally.
func BestPracticeSettings() []Setting {
	return []Setting{
		// Branch & Init (3)
		{Key: "init.defaultBranch", Value: "main", Category: "Branch & Init", Description: "modern default branch name"},
		{Key: "branch.sort", Value: "-committerdate", Category: "Branch & Init", Description: "sort branches by most recently used"},
		{Key: "column.ui", Value: "auto", Category: "Branch & Init", Description: "display branch list in columns"},

		// Pull & Push (5)
		{Key: "pull.rebase", Value: "true", Category: "Pull & Push", Description: "avoid accidental merge commits on pull"},
		{Key: "push.default", Value: "simple", Category: "Pull & Push", Description: "push current branch to same-named remote"},
		{Key: "push.autoSetupRemote", Value: "true", Category: "Pull & Push", Description: "auto-configure upstream on first push"},
		{Key: "fetch.prune", Value: "true", Category: "Pull & Push", Description: "remove stale remote-tracking branches"},
		{Key: "fetch.pruneTags", Value: "true", Category: "Pull & Push", Description: "remove stale remote-tracking tags"},

		// Diff & Merge (3)
		{Key: "diff.algorithm", Value: "histogram", Category: "Diff & Merge", Description: "more readable diffs for moved/refactored code"},
		{Key: "diff.colorMoved", Value: "plain", Category: "Diff & Merge", Description: "highlight moved code in a distinct color"},
		{Key: "merge.conflictstyle", Value: "zdiff3", Category: "Diff & Merge", Description: "show original base in conflict markers"},

		// Rebase (2)
		{Key: "rebase.autoSquash", Value: "true", Category: "Rebase", Description: "auto-combine fixup! commits"},
		{Key: "rebase.autoStash", Value: "true", Category: "Rebase", Description: "stash/restore dirty working tree around rebase"},

		// Workflow (3)
		{Key: "commit.verbose", Value: "true", Category: "Workflow", Description: "show diff while writing commit message"},
		{Key: "rerere.enabled", Value: "true", Category: "Workflow", Description: "remember and reuse conflict resolutions"},
		{Key: "help.autocorrect", Value: "prompt", Category: "Workflow", Description: "ask before auto-correcting mistyped commands"},

		// Performance (2)
		{Key: "core.fsmonitor", Value: "true", Category: "Performance", Description: "faster git status via filesystem events"},
		{Key: "core.untrackedCache", Value: "true", Category: "Performance", Description: "cache untracked file info for speed"},
	}
}

// EditorValue maps editor shorthand to the git core.editor value.
func EditorValue(editor string) string {
	switch editor {
	case "vscode":
		return "code --wait"
	case "vim":
		return "vim"
	case "nano":
		return "nano"
	case "emacs":
		return "emacs"
	default:
		return editor
	}
}
