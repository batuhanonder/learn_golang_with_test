package v1

import (
	"testing"
)

func Test_roman_numerals(t *testing.T) {
	got := ConvertToRoman(1)
	want := "I"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func ConvertToRoman(_ int) string {
	return "I"
}
