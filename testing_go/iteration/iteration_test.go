package iteration

import "testing"

func TestFor(t *testing.T) {
	repeat := ForLoop("a")
	expect := "aaaaa"

	if repeat != expect {
		t.Errorf("Expected %q got %q", expect, repeat)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ForLoop("a")
	}
}
