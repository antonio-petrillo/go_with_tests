package numeral

import (
	"fmt"
	"testing"
	"testing/quick"
)

// func TestRomanNumerals(t *testing.T) {
// 	t.Run("1 gets converted to I", func(t *testing.T) {
// 		got := ConvertToRoman(1)
// 		want := "I"

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})

// 	t.Run("2 gets converted to II", func(t *testing.T) {
// 		got := ConvertToRoman(2)
// 		want := "II"

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})
// }

// func TestRomanNumerals(t *testing.T) {
// 	cases := []struct{
// 		Description string
// 		Arabic int
// 		Want string
// 	} {
// 		{"1 gets converted to I", 1, "I"},
// 		{"2 gets converted to II", 2, "II"},
// 		{"3 gets converted to III", 3, "III"},
// 		{"4 gets converted to IV", 4, "IV"},
// 		{"5 gets converted to V", 5, "V"},
// 		{"6 gets converted to VI", 6, "VI"},
// 		{"7 gets converted to VII", 7, "VII"},
// 		{"8 gets converted to VIII", 8, "VIII"},
// 		{"9 gets converted to IX", 9, "IX"},
// 		{"10 gets converted to X", 10, "X"},
// 		{"40 gets converted to XL", 40, "XL"},
// 		{"47 gets converted to XLVII", 47, "XLVII"},
// 		{"49 gets converted to XLIX", 49, "XLIX"},
// 		{"50 gets converted to L", 50, "L"},
// 		{"100 gets converted to C", 100, "C"},
// 		{"500 gets converted to D", 500, "D"},
// 		{"1000 gets converted to M", 1000, "M"},
// 	}

// 	t.Run("arabic to roman", func (t *testing.T) {
// 		for _, test := range cases {
// 			t.Run(test.Description, func(t *testing.T) {
// 				got := ConvertToRoman(test.Arabic)
// 				if got != test.Want {
// 					t.Errorf("got %q, want %q", got, test.Want)
// 				}
// 			})
// 		}
// 	})
// }

func TestRomanConversion(t *testing.T) {
	cases := []struct{
		Arabic int
		Roman string
	} {
		{Arabic: 1, Roman: "I"},
			{Arabic: 2, Roman: "II"},
			{Arabic: 3, Roman: "III"},
			{Arabic: 4, Roman: "IV"},
			{Arabic: 5, Roman: "V"},
			{Arabic: 6, Roman: "VI"},
			{Arabic: 7, Roman: "VII"},
			{Arabic: 8, Roman: "VIII"},
			{Arabic: 9, Roman: "IX"},
			{Arabic: 10, Roman: "X"},
			{Arabic: 40, Roman: "XL"},
			{Arabic: 47, Roman: "XLVII"},
			{Arabic: 49, Roman: "XLIX"},
			{Arabic: 50, Roman: "L"},
			{Arabic: 100, Roman: "C"},
			{Arabic: 500, Roman: "D"},
			{Arabic: 900, Roman: "CM"},
			{Arabic: 1000, Roman: "M"},
			{Arabic: 1984, Roman: "MCMLXXXIV"},
			{Arabic: 3999, Roman: "MMMCMXCIX"},
			{Arabic: 2014, Roman: "MMXIV"},
			{Arabic: 1006, Roman: "MVI"},
			{Arabic: 798, Roman: "DCCXCVIII"},
		}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func (t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}

}

func TestConvertingToArabic(t *testing.T) {
	cases := []struct{
		Arabic int
		Roman string
	} {
		{Arabic: 1, Roman: "I"},
			{Arabic: 2, Roman: "II"},
			{Arabic: 3, Roman: "III"},
			{Arabic: 4, Roman: "IV"},
			{Arabic: 5, Roman: "V"},
			{Arabic: 6, Roman: "VI"},
			{Arabic: 7, Roman: "VII"},
			{Arabic: 8, Roman: "VIII"},
			{Arabic: 9, Roman: "IX"},
			{Arabic: 10, Roman: "X"},
			{Arabic: 40, Roman: "XL"},
			{Arabic: 47, Roman: "XLVII"},
			{Arabic: 49, Roman: "XLIX"},
			{Arabic: 50, Roman: "L"},
			{Arabic: 100, Roman: "C"},
			{Arabic: 500, Roman: "D"},
			{Arabic: 900, Roman: "CM"},
			{Arabic: 1000, Roman: "M"},
			{Arabic: 1984, Roman: "MCMLXXXIV"},
			{Arabic: 3999, Roman: "MMMCMXCIX"},
			{Arabic: 2014, Roman: "MMXIV"},
			{Arabic: 1006, Roman: "MVI"},
			{Arabic: 798, Roman: "DCCXCVIII"},
		}

	for _, test := range cases[:2] {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func (t *testing.T){
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
 		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic int) bool {
		if (arabic < 0 || arabic > 3999){
			return true
		}
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
