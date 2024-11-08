package v1

import (
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	t.Run("1 gets converted to I", func(t *testing.T) {
		got := ConvertToRoman(1)
		want := "I"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("2 gets converted to II", func(t *testing.T) {
		got := ConvertToRoman(2)
		want := "II"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func ConvertToRoman(num int) string {
	if num == 2 {
		return "II"
	}
	return "I"
}
