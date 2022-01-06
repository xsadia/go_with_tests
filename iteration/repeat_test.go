package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "AAAAA"

	if repeated != expected {
		t.Errorf("expected %q, got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	val := Repeat("a", 5)

	fmt.Println(val)
	// Output: AAAAA
}
