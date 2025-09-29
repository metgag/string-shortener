package main

import (
	"fmt"
	"log"
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

	fmt.Println(1)
	fmt.Println(StringShortener(v))
	fmt.Println()

	fmt.Println(2)
	fmt.Println(StringShortener(w))
	fmt.Println()

	fmt.Println(3)
	fmt.Println(StringShortener(y))
	fmt.Println()

	fmt.Println(4)
	fmt.Println(StringShortener(x))
	fmt.Println()

	fmt.Println(5)
	fmt.Println(StringShortener(z))
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

	for i := range inputStr {
		if len(inputStr) <= 1 {
			break
		}
		splitStr := strings.Split(inputStr, "")

		if splitStr[i] == splitStr[i+1] {
			inputStr = strings.Replace(
				inputStr,
				strings.Join([]string{splitStr[i], splitStr[i+1]}, ""),
				"",
				1,
			)
			log.Println(inputStr, "========> len", len(inputStr))

			i = 0
		}
	}

	return inputStr
}
