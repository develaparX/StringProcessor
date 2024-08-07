package main

import (
	"fmt"
	"skkmigas/lettergroup"
)

func main() {
	input := []string{"Abc", "bCd"}
	result := lettergroup.GroupAndSortLetters(input)
	fmt.Println(result) // Output: bACcd

	input = []string{"Oke", "One"}
	result = lettergroup.GroupAndSortLetters(input)
	fmt.Println(result) // Output: Oekn

	input = []string{"Pendanaan", "Terproteksi", "Untuk", "Dampak", "Berarti"}
	result = lettergroup.GroupAndSortLetters(input)
	fmt.Println(result) // Output: aenrktiBDPTUdimu
}
