package v1

import "strings"

// ConvertToInt converts a Roman Numeral to a num number.
func ConvertToInt(roman string) uint16 {
	var num uint16 = 0

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			num += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return num
}

// ConvertToRoman converts an num number to a Roman Numeral.
func ConvertToRoman(num uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for num >= numeral.Value {
			result.WriteString(numeral.Symbol)
			num -= numeral.Value
		}
	}

	return result.String()
}

type romanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []romanNumeral{
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
