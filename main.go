package main

import (
	"fmt"

	"github.com/deenski/goStringUtils"
)

func main() {
	fmt.Println(goStringUtils.Reverse("!oG ,olleH"))

	fmt.Println(goStringUtils.ToLower("ALL OF THESE LETTERS ARE LOWER CASE"))

	fmt.Println(goStringUtils.ToUpper("and this string is upper case"))

	fmt.Print(goStringUtils.CompareThese("four", "five"))
	fmt.Print(goStringUtils.CompareThese("two", "three"))
	fmt.Print(goStringUtils.CompareThese("nineteen", "one"))

	a := "This is a "
	b := "concatinated string."

	fmt.Println(goStringUtils.ConcatThese(a, b))
}
