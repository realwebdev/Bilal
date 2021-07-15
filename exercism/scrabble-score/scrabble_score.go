package scrabble

import "strings"

// func Scrabble(str string) int {

// 	var vale = map[rune]int{
// 		'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1,
// 		'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
// 		'D': 2, 'G': 2,
// 		'B': 3, 'C': 3, 'M': 3, 'P': 3,
// 		'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
// 		'K': 5,
// 		'J': 8, 'X': 8,
// 		'Q': 10, 'Z': 10,
// 	}

// 	str = strings.ToUpper(str)

// 	sum := 0
// 	for _, values := range str {

// 		sum += vale[values]

// 	}

var letterValues = map[rune]int{
	'a': 1, 'b': 3, 'c': 3, 'd': 2, 'e': 1,
	'f': 4, 'g': 2, 'h': 4, 'i': 1, 'j': 8,
	'k': 5, 'l': 1, 'm': 3, 'n': 1, 'o': 1,
	'p': 3, 'q': 10, 'r': 1, 's': 1, 't': 1,
	'u': 1, 'v': 4, 'w': 4, 'x': 8, 'y': 4,
	'z': 10,
}

// Score computes the number of points that a word is worth.
func Score(word string) int {
	word = strings.ToLower(word)

	sum := 0
	for _, letter := range word {
		sum += letterValues[letter]
	}
	return sum
}
