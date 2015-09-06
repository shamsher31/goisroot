package root

import "testing"

func TestIs(t *testing.T) {
	if Is() == true {
		t.Fatal("Should return true")
	}
}
