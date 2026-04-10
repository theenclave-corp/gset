package prompt_test

import (
	"testing"

	"github.com/theenclave-corp/gset/internal/prompt"
)

// When all fields are pre-filled, AskIdentity must return nil without
// launching any interactive form (which would hang in a non-TTY environment).
func TestAskIdentity_SkipsWhenPreFilled(t *testing.T) {
	id := &prompt.Identity{
		Name:   "Alice",
		Email:  "alice@example.com",
		Editor: "vim",
	}
	if err := prompt.AskIdentity(id, true); err != nil {
		t.Fatalf("AskIdentity with pre-filled fields returned error: %v", err)
	}
	// Values must be unchanged
	if id.Name != "Alice" {
		t.Errorf("Name changed: got %q", id.Name)
	}
	if id.Editor != "vim" {
		t.Errorf("Editor changed: got %q", id.Editor)
	}
}

func TestAskIdentity_NoEditorWhenWithEditorFalse(t *testing.T) {
	// withEditor=false + pre-filled name/email = no fields to prompt, returns nil
	id := &prompt.Identity{
		Name:  "Bob",
		Email: "bob@example.com",
	}
	if err := prompt.AskIdentity(id, false); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id.Editor != "" {
		t.Errorf("Editor should be empty when withEditor=false, got %q", id.Editor)
	}
}
