package main

import "fmt"

func intToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			roman += syms[i]
			num -= val[i]
		}
	}
	return roman
}

func main() {
	fmt.Println(intToRoman(4))    // IV
	fmt.Println(intToRoman(9))    // IX
	fmt.Println(intToRoman(23))   // XXIII
	fmt.Println(intToRoman(2021)) // MMXXI
	fmt.Println(intToRoman(1646)) // MDCXLVI
}
