package config_test

import (
	"testing"

	"github.com/theenclave-corp/gset/internal/config"
)

func TestBestPracticeSettings_Count(t *testing.T) {
	const want = 18
	got := len(config.BestPracticeSettings())
	if got != want {
		t.Errorf("BestPracticeSettings() returned %d settings, want %d", got, want)
	}
}

func TestBestPracticeSettings_HaveRequiredFields(t *testing.T) {
	for _, s := range config.BestPracticeSettings() {
		if s.Key == "" {
			t.Errorf("setting has empty Key: %+v", s)
		}
		if s.Value == "" {
			t.Errorf("setting %q has empty Value", s.Key)
		}
		if s.Category == "" {
			t.Errorf("setting %q has empty Category", s.Key)
		}
		if s.Description == "" {
			t.Errorf("setting %q has empty Description", s.Key)
		}
	}
}

func TestEditorValue(t *testing.T) {
	cases := []struct {
		editor string
		want   string
	}{
		{"vscode", "code --wait"},
		{"vim", "vim"},
		{"nano", "nano"},
		{"emacs", "emacs"},
		{"unknown", "unknown"}, // passthrough
	}
	for _, c := range cases {
		got := config.EditorValue(c.editor)
		if got != c.want {
			t.Errorf("EditorValue(%q) = %q, want %q", c.editor, got, c.want)
		}
	}
}
