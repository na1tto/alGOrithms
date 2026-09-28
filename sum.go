package module01

import "fmt"

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	var tmp int
	for _, i := range numbers {
		tmp += i
		fmt.Print(tmp)
	}
	return tmp
}
