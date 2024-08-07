package lettergroup

import (
	"sort"
	"strings"
)

// GroupAndSortLetters groups each letter (case-sensitive) and sorts them by frequency and alphabetical order.
func GroupAndSortLetters(input []string) string {
	letterCount := make(map[rune]int)

	// Count occurrences of each letter
	for _, word := range input {
		for _, letter := range word {
			letterCount[letter]++
		}
	}

	// Create a slice of letters for sorting
	type letterFrequency struct {
		letter    rune
		frequency int
	}
	var letters []letterFrequency
	for letter, frequency := range letterCount {
		letters = append(letters, letterFrequency{letter, frequency})
	}

	// Sort the letters by frequency (desc), then by letter (capital letters first, then lowercase)
	sort.Slice(letters, func(i, j int) bool {
		if letters[i].frequency == letters[j].frequency {
			return letters[i].letter < letters[j].letter
		}
		return letters[i].frequency > letters[j].frequency
	})

	// Construct the result string
	var result strings.Builder
	for _, lf := range letters {
		result.WriteRune(lf.letter)
	}

	return result.String()
}
