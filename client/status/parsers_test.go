package status

import (
	"errors"
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
)

func TestFromAuthStatus(t *testing.T) {
	tests := []struct {
		name     string
		st       auth.Status
		wantNil  bool
		wantType string
	}{
		{"NoAuth", auth.NoAuth, false, "*status.ErrNoAuth"},
		{"Ok", auth.Ok, true, ""},
		{"OTP", auth.OTP, false, "*status.ErrMissingOTPAuth"},
		{"Unknown", auth.Unknown, false, "*status.ErrUnknownAuth"},
		{"Other", auth.Other, false, "*status.ErrAuthFailed"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fromAuthStatus(tt.st)
			if tt.wantNil {
				if err != nil {
					t.Fatalf("expected nil error, got %T: %v", err, err)
				}
				return
			}
			if err == nil {
				t.Fatalf("expected non-nil error for %s", tt.name)
			}
			if got := fmtType(err); got != tt.wantType {
				t.Fatalf("expected type %s, got %s", tt.wantType, got)
			}
		})
	}
}

func TestFilterIssues(t *testing.T) {
	in := []string{"Incorrect", "foo", "The token supplied is invalid.", "bar", "You may not reuse tokens.", "baz"}

	remaining, errs := filterIssues(append([]string(nil), in...)) // copy

	// 'Incorrect' and 'The token supplied is invalid.' are known and produce errors
	if len(errs) != 2 {
		t.Fatalf("expected 2 errors, got %d", len(errs))
	}

	// Remaining should preserve order of unknowns
	wantRemaining := []string{"foo", "bar", "baz"}
	if len(remaining) != len(wantRemaining) {
		t.Fatalf("unexpected remaining length: %d", len(remaining))
	}
	for i := range wantRemaining {
		if remaining[i] != wantRemaining[i] {
			t.Fatalf("remaining[%d] expected %q, got %q", i, wantRemaining[i], remaining[i])
		}
	}

	// Ensure we received expected error types
	var sawAuthFailed, sawOTPFailed bool
	var vAuth *ErrAuthFailed
	var vOTP *ErrOTPAuthFailed
	for _, e := range errs {
		if errors.As(e, &vAuth) {
			sawAuthFailed = true
		}
		if errors.As(e, &vOTP) {
			sawOTPFailed = true
		}
	}
	if !sawAuthFailed || !sawOTPFailed {
		t.Fatalf("expected both ErrAuthFailed and ErrOTPAuthFailed in errs; got %+v", errs)
	}
}

func TestFromIssues(t *testing.T) {
	in := []string{"alpha", "beta"}
	errs := fromIssues(append([]string(nil), in...))

	if len(errs) != 2 {
		t.Fatalf("expected 2 errors, got %d", len(errs))
	}
	var vErr *ErrHeNet
	for i, e := range errs {
		if !errors.As(e, &vErr) {
			t.Fatalf("expected ErrHeNet, got %T", e)
		}
		if e.Error() != in[i] {
			t.Fatalf("expected error message %q, got %q", in[i], e.Error())
		}
	}
}

// fmtType returns the concrete type string of an error (e.g. "*status.ErrNoAuth").
func fmtType(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("%T", err)
}
