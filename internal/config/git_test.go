package config_test

import (
	"os"
	"testing"

	"github.com/theenclave-corp/gset/internal/config"
)

// tempScope returns scope args pointing at a fresh temp file.
func tempScope(t *testing.T) []string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "gitconfig-*")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	return []string{"--file", f.Name()}
}

func TestSetAndGet(t *testing.T) {
	scope := tempScope(t)

	if err := config.Set(scope, "user.name", "Test User"); err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	val, err := config.Get(scope, "user.name")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if val != "Test User" {
		t.Errorf("Get = %q, want %q", val, "Test User")
	}
}

func TestGetMissing(t *testing.T) {
	scope := tempScope(t)

	_, err := config.Get(scope, "user.name")
	if err == nil {
		t.Fatal("expected error for missing key, got nil")
	}
}

func TestList(t *testing.T) {
	scope := tempScope(t)

	pairs := map[string]string{
		"user.name":  "Alice",
		"user.email": "alice@example.com",
	}
	for k, v := range pairs {
		if err := config.Set(scope, k, v); err != nil {
			t.Fatalf("Set(%q) failed: %v", k, err)
		}
	}

	all, err := config.List(scope)
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}
	for k, want := range pairs {
		if got := all[k]; got != want {
			t.Errorf("List[%q] = %q, want %q", k, got, want)
		}
	}
}
