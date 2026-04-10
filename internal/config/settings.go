package config

// Setting represents a single git config key/value pair with metadata.
type Setting struct {
	Key         string
	Value       string
	Category    string
	Description string
}

// EditorValue maps a human-readable editor name to the git config value.
func EditorValue(editor string) string {
	return editor
}

// BestPracticeSettings returns the curated list of best-practice git settings.
func BestPracticeSettings() []Setting {
	return nil
}
