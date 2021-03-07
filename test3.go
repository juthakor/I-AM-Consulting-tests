package main

import (
	"fmt"
)

var resultArr []string

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func initiator(str string) {
	r := []rune(str)
	Perm(r, func(a []rune) {
		if !contains(resultArr, string(a)) {
			resultArr = append(resultArr, string(a))
		}
	})
	fmt.Println(resultArr)
}

func main() {
	initiator("ABC")
}
