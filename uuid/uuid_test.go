package uuid

import "testing"

func TestNew(t *testing.T) {
	s := New()
	if s == "" {
		t.Errorf("New() should not be empty")
	}
}
