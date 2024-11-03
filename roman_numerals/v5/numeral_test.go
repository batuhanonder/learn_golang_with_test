package v1

import (
	"fmt"
	"testing"
	"testing/quick"
)

var (
	cases = []struct {
		num   uint16
		Roman string
	}{
		{num: 1, Roman: "I"},
		{num: 2, Roman: "II"},
		{num: 3, Roman: "III"},
		{num: 4, Roman: "IV"},
		{num: 5, Roman: "V"},
		{num: 6, Roman: "VI"},
		{num: 7, Roman: "VII"},
		{num: 8, Roman: "VIII"},
		{num: 9, Roman: "IX"},
		{num: 10, Roman: "X"},
		{num: 14, Roman: "XIV"},
		{num: 18, Roman: "XVIII"},
		{num: 20, Roman: "XX"},
		{num: 39, Roman: "XXXIX"},
		{num: 40, Roman: "XL"},
		{num: 47, Roman: "XLVII"},
		{num: 49, Roman: "XLIX"},
		{num: 50, Roman: "L"},
		{num: 100, Roman: "C"},
		{num: 90, Roman: "XC"},
		{num: 400, Roman: "CD"},
		{num: 500, Roman: "D"},
		{num: 900, Roman: "CM"},
		{num: 1000, Roman: "M"},
		{num: 1984, Roman: "MCMLXXXIV"},
		{num: 3999, Roman: "MMMCMXCIX"},
		{num: 2014, Roman: "MMXIV"},
		{num: 1006, Roman: "MVI"},
		{num: 798, Roman: "DCCXCVIII"},
	}
)

func TestConvertingToRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to '%s", test.num, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.num)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertingToInt(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.num), func(t *testing.T) {
			got := ConvertToInt(test.Roman)
			if got != test.num {
				t.Errorf("got %d, want %d", got, test.num)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(num uint16) bool {
		if num > 3999 {
			return true
		}
		t.Log("testing", num)
		roman := ConvertToRoman(num)
		fromRoman := ConvertToInt(roman)
		return fromRoman == num
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
