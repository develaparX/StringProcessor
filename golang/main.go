package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type CharFreq struct {
	Char rune
	Freq int
}

func processStrings(inputArray []string) string {
	// Gabungkan menjadi satu string
	combinedString := strings.Join(inputArray, "")

	// Inisialisasi map untuk menghitung frekuensi huruf
	frequency := make(map[rune]int)

	// Hitung frekuensi huruf
	for _, char := range combinedString {
		frequency[char]++
	}

	// Buat slice dari karakter dan frekuensinya
	var charArray []CharFreq
	for char, freq := range frequency {
		charArray = append(charArray, CharFreq{Char: char, Freq: freq})
	}

	// Urutkan berdasarkan frekuensi dan aturan tambahan
	sort.Slice(charArray, func(i, j int) bool {
		if charArray[i].Freq == charArray[j].Freq {
			// Jika frekuensi sama, urutkan berdasarkan aturan tambahan
			if strings.ToLower(string(charArray[i].Char)) == strings.ToLower(string(charArray[j].Char)) {
				// Jika karakter yang diabaikan case-nya sama, urutkan berdasarkan nilai ASCII
				return charArray[i].Char < charArray[j].Char
			}
			if unicode.IsUpper(charArray[i].Char) && unicode.IsLower(charArray[j].Char) {
				// Jika karakter pertama uppercase dan karakter kedua lowercase, karakter pertama lebih kecil
				return true
			}
			if unicode.IsLower(charArray[i].Char) && unicode.IsUpper(charArray[j].Char) {
				// Jika karakter pertama lowercase dan karakter kedua uppercase, karakter pertama lebih besar
				return false
			}
			// Urutkan berdasarkan karakter case-insensitive
			return strings.ToLower(string(charArray[i].Char)) < strings.ToLower(string(charArray[j].Char))
		}
		// Urutkan berdasarkan frekuensi dari besar ke kecil
		return charArray[i].Freq > charArray[j].Freq
	})

	// Buat string hasil akhir dengan hanya satu kali kemunculan setiap karakter
	var result strings.Builder
	seen := make(map[rune]bool)

	for _, item := range charArray {
		if !seen[item.Char] {
			result.WriteRune(item.Char)
			seen[item.Char] = true
		}
	}

	return result.String()
}

func main() {
	inputArray := []string{"Pendanaan", "Terproteksi", "Untuk", "Dampak", "Berarti"}
	output := processStrings(inputArray)
	fmt.Println(output)
}
