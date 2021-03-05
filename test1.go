package main

import (
	"fmt"
	"strings"
	"math"
)

func isPalindrom(word string) bool {
    word = strings.ToLower(word)
    for i := 0; float64(i) < math.Floor(float64(len(word) / 2)); i++ {
        if (word[i] != word[len(word) - i - 1]) {
            return false
        }
    }
    return true
}

func main() {
	word := "Deleveled"
	fmt.Println(word, isPalindrom(word))
}
