package propertytest

import "strings"

type RomanNumberal struct {
	Value  int
	Symbol string
}

var RomanNumberals = []RomanNumberal{
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

func ConvertToRoman(arabic int) string {
	for _, numberal := range RomanNumberals {
		for arabic >= numberal.Value {
			return numberal.Symbol + ConvertToRoman(arabic-numberal.Value)
		}
	}
	return ""
}

func ConvertToArabic(roman string) int {
	for _, numberal := range RomanNumberals {
		if strings.HasPrefix(roman, numberal.Symbol) {
			return numberal.Value + ConvertToArabic(roman[len(numberal.Symbol):])
		}
	}
	return 0
}
