package iteration

import "testing"

func Repeat(character string) string {
	var repeat string
	for i := 0; i < 5; i++ {
		repeat = repeat + character
	} 
	return repeat
}
func TestRepeat(t *testing.T) {
	repeat := Repeat("a")
	expected := "aaaaa"

	if repeat != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeat)
	}
}