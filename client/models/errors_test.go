package models

import "testing"

var _ error = (*ErrFormat)(nil)

func TestErrFormat_Error_WithField(t *testing.T) {
	e := &ErrFormat{Field: "name", Reason: "is required"}
	got := e.Error()
	want := `invalid "name", is required`
	if got != want {
		t.Fatalf("Error() = %q, want %q", got, want)
	}
}

func TestErrFormat_Error_NoField(t *testing.T) {
	e := &ErrFormat{Reason: "missing value"}
	got := e.Error()
	want := "invalid record, missing value"
	if got != want {
		t.Fatalf("Error() = %q, want %q", got, want)
	}
}
