package root

import "testing"

// go test -v
func TestIs(t *testing.T) {
	if Is() == true {
		t.Fatal("Should return true")
	}
}

// go test -test.bench=".*"
func BenchmarkIs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Is()
	}
}
