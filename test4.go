package main

import (
	"fmt"
	"math"
)

func findSum(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += int(math.Pow(float64(i), float64(i)))
    }
    return sum
}

func main() {
	fmt.Println(findSum(3))
}
