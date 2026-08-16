package rbac

import "testing"

func TestProbeEmptyParentIsRejected(t *testing.T) {
	s := New()
	if err := s.PutRole(Role{ID: "reader", Parents: []string{""}}); err == nil {
		t.Fatal("empty parent must be rejected")
	}
	if s.HasRole("reader") {
		t.Fatal("rejected role must not be stored")
	}
}
