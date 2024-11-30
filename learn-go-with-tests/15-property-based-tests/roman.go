package roman

import "strings"

type Numeral struct {
	Value  uint16
	Symbol string
}

//nolint:gochecknoglobals
var allRomanNumerals = []Numeral{
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

func ConvertToRoman(arabic uint16) string {
	result := strings.Builder{}

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

const lengthOfSubtractSymbol = 2
const lengthOfSymbol = 1

func ConvertToArabic(roman string) uint16 {
	var result uint16

	for len(roman) > 0 {
		parsed := false

		if len(roman) >= lengthOfSubtractSymbol {
			for _, numeral := range allRomanNumerals {
				if roman[0:lengthOfSubtractSymbol] == numeral.Symbol {
					result += numeral.Value
					roman = roman[lengthOfSubtractSymbol:]
					parsed = true

					break
				}
			}
		}

		if !parsed && len(roman) >= lengthOfSymbol {
			if !parsed {
				for _, numeral := range allRomanNumerals {
					if roman[0:lengthOfSymbol] == numeral.Symbol {
						result += numeral.Value
						roman = roman[lengthOfSymbol:]
						parsed = true

						break
					}
				}
			}
		}

		if !parsed {
			panic("unexpected")
		}
	}

	return result
}
