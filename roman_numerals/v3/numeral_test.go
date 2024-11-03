package v1

import (
	"fmt"
	"strings"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		num   int
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
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.num, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.num)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(num int) string {

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for num >= numeral.Value {
			result.WriteString(numeral.Symbol)
			num -= numeral.Value
		}
	}

	return result.String()
}
