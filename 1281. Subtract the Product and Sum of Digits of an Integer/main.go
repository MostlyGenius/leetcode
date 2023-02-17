package main

import "fmt"

func subtractProductAndSum(n int) int {
	i := 0
	prodResult := 1
	sumResult := 0

	for i < n {
		prodResult = prodResult * n[i]
	}

	for i < len(n) {
		sumResult = sumResult + n[i]
	}

	return prodResult - sumResult
}

func main() {
	n := 234
	testCase := subtractProductAndSum(n)
	fmt.Println(testCase)
}
