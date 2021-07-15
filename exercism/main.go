package main

import (
	"fmt"
)

func main() {

	var rev_string string = "cool"
	var outp string

	fmt.Scanln(rev_string)

	for i := len(rev_string); i < len(rev_string); i-- {
		outp := copy(outp[i], rev_string[i-1])

		rev_string = outp
	}

	fmt.Println(rev_string)

}


package bob

import "strings"

// Hey returns Bob's responses to a given remark.
func Hey(remark string) string {
	switch remark = strings.TrimSpace(remark); {
	case silent(remark):
		return "Fine. Be that way!"
	case yelling(remark):
		if asking(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	case asking(remark):
		return "Sure."
	default:
		return "Whatever."
	}
}

func yelling(remark string) bool {
	return strings.ToUpper(remark) == remark && strings.ToLower(remark) != strings.ToUpper(remark)
}

func asking(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func silent(remark string) bool {
	return remark == ""
}




import "math"

var allergens = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

func Allergies(i uint) []string {
	var result []string
	for _, v := range allergens {
		if AllergicTo(i, v) {
			result = append(result, v)
		}
	}
	return result
}

func AllergicTo(i uint, allergen string) bool {
	index := indexOf(allergens, allergen)
	x := uint(math.Pow(2.0, float64(index)))
	return (i & x) > 0
}

func indexOf(slice []string, s string) int {
	for p, v := range slice {
		if v == s {
			return p
		}
	}
	return -1
}




var allergens = map[string]uint{"eggs": 1, "peanuts": 2, "shellfish": 4, "strawberries": 8, "tomatoes": 16, "chocolate": 32, "pollen": 64, "cats": 128}

func Allergies(v uint, i string) []string {
	var result []string
	for _, d := range allergens {
		if d == v {
			result = append(result, i)
		}
	}
	return result
}
