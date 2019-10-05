package main

import (
	"fmt"

	"github.com/deenski/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))

	fmt.Println(stringutil.ToLower("ALL OF THESE LETTERS ARE LOWER CASE"))

	fmt.Println(stringutil.ToUpper("and this string is upper case"))

	fmt.Print(stringutil.CompareThese("four", "five"))
	fmt.Print(stringutil.CompareThese("two", "three"))
	fmt.Print(stringutil.CompareThese("nineteen", "one"))

	a := "This is a "
	b := "concatinated string."

	fmt.Println(stringutil.ConcatThese(a, b))
}
