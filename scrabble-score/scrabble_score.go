package scrabble

import "unicode"

// Answer using a map of rune to int
// func Score(word string) int {
// 	score := map[rune]int{
// 		'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
// 		'D': 2, 'G': 2,
// 		'B': 3, 'C': 3, 'M': 3, 'P': 3,
// 		'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
// 		'K': 5,
// 		'J': 8, 'X': 8,
// 		'Q': 10, 'Z': 10,
// 	}
// 	total := 0
// 	for _, r := range word {
// 		if val, ok := score[rune(unicode.ToUpper(r))]; ok {
// 			total += val
// 		}
// 	}
// 	return total
// }

// Alternative answer using a switch statement
func Score(word string) int {
	score := 0
	for _, letter := range word {
		switch unicode.ToUpper(letter) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score += 1
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J', 'X':
			score += 8
		case 'Q', 'Z':
			score += 10
		default:
		}
	}
	return score
}
