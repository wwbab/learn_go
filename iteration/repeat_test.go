package iteration

import (
	"strings"
	"testing"
)

const repeatCount = 5

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	} 
	return repeated.String()
}
func TestRepeat(t *testing.T) {
	repeat := Repeat("a")
	expected := "aaaaa"

	if repeat != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeat)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		Repeat("a")
	}
}