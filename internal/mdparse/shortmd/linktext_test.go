package shortmd

import "testing"

func TestLinkText(t *testing.T) {
	const md = `array of [role](#DOCS_TOPICS_PERMISSIONS/role-object) object ids`

	text := LinkText(md)

	if text != "role" {
		t.Fatalf("expected %q, got %q", "role", text)
	}
}
