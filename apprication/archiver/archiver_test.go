package archiver

import "testing"

func TestArchive(t *testing.T) {
	foo := Archive()

	if foo != "Hello world" {
		t.Error()
	}
}
