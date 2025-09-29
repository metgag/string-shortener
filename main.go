package main

import (
	"fmt"
	"strings"
)

func main() {
	// foo := []int{0, 2, 1, 5, 3, 4}
	// bar := []int{5, 0, 1, 2, 3, 4}

	// fmt.Println(BuildArray(foo))
	// fmt.Println(BuildArray(bar))

	// string shortener
	v := "aab"
	/*
		b
	*/

	w := "aaabba"
	/*
		abba
		aa
		""
	*/

	x := "aabb"
	/*
		bb
		""
	*/

	y := "aaa"
	/*
		a
	*/

	z := "aaabccddd"
	/*
		abccddd
		abddd
		abd
	*/

	fmt.Println("1.", StringShortener(v))
	fmt.Println()

	fmt.Println("2.", StringShortener(w))
	fmt.Println()

	fmt.Println("3.", StringShortener(x))
	fmt.Println()

	fmt.Println("4.", StringShortener(y))
	fmt.Println()

	fmt.Println("5.", StringShortener(z))
}

/*
input: nums []int
output: ans [len(nums)]int
process: create new arr from nums start from nums[i]
*/
func BuildArray(nums []int) []int {
	ans := make([]int, len(nums))

	for i := range nums {
		ans = append(ans, nums[nums[i]])
	}

	return ans
}

func StringShortener(inputStr string) string {
	if inputStr == "" {
		return ""
	}

	for i := 0; i < len(inputStr)-1; i++ {
		if inputStr[i] == inputStr[i+1] {
			inputStr = strings.Replace(
				inputStr,
				strings.Join([]string{
					string(inputStr[i]),
					string(inputStr[i+1]),
				}, ""),
				"",
				1,
			)

			i = 0
		}
	}
	if len(inputStr) > 1 {
		if inputStr[0] == inputStr[1] {
			return "empty string"
		}
	}

	return inputStr
}
