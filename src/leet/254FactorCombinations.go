/*
Numbers can be regarded as product of its factors. For example,

8 = 2 x 2 x 2;
  = 2 x 4.
Write a function that takes an integer n and return all possible combinations of its factors.

Note:
Each combination's factors must be sorted ascending, for example: The factors of 2 and 6 is [2, 6], not [6, 2].
You may assume that n is always positive.
Factors should be greater than 1 and less than n.
Examples:
input: 1
output:
[]
input: 37
output:
[]
input: 12
output:
[
  [2, 6],
  [2, 2, 3],
  [3, 4]
]
input: 32
output:
[
  [2, 16],
  [2, 2, 8],
  [2, 2, 2, 4],
  [2, 2, 2, 2, 2],
  [2, 4, 4],
  [4, 8]
]
*/

package main

import "fmt"

func main() {
	fact12 := getFactors(12)
	fmt.Print("Factors for 12:\n")
	printAnswer(fact12)

	fact1 := getFactors(1)
	fmt.Print("Factors for 1:\n")
	printAnswer(fact1)

	fact37 := getFactors(37)
	fmt.Print("Factors for 37:\n")
	printAnswer(fact37)

	fact32 := getFactors(32)
	fmt.Print("Factors for 32:\n")
	printAnswer(fact32)

}

func printAnswer(answer [][]int) {
	for i:=0; i<len(answer); i++ {
		fmt.Printf("%v\n", answer[i])
	}
}

func appendToAll(subResult [][]int, tail []int) {
	for i:=0; i<len(subResult); i++ {
		subResult[i] = append(subResult[i], tail...)
	}
}


func internalGetFactors(n, maxFactor int, allowOne bool) [][]int {
	result := make([][]int, 0)

	if n==1 {
		return result
	}

	if allowOne && n <= maxFactor {
		result = append(result, []int { n })
	}

	for f:=maxFactor; f>1; f-- {
		if n % f != 0 {
			continue
		}

		tail := []int { f }
		newN := n / f

		subResult := internalGetFactors(newN, f - 1, true)
		appendToAll(subResult, tail)
		// merge results
		result = append(result, subResult...)


		for newN % f == 0 {
			tail = append(tail, f)
			newN = newN / f

			if newN != 1 {
				subResult = internalGetFactors(newN, f - 1, true)
				appendToAll(subResult, tail)
				// merge results
				result = append(result, subResult...)
			} else {
				result = append(result, tail)
			}


		}
	}

	return result
}

func getFactors(n int) [][]int {
	return internalGetFactors(n, n-1, false)
}
