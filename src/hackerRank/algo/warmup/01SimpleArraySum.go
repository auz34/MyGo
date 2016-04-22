/*
Given an array of NN integers, can you find the sum of its elements?

Input Format

The first line contains an integer, NN, denoting the size of the array.
The second line contains NN space-separated integers representing the array's elements.

Output Format

Print the sum of the array's elements as a single integer.

Sample Input

6
1 2 3 4 10 11
Sample Output

31
Explanation

We print the sum of the array's elements, which is: 1+2+3+4+10+11=311+2+3+4+10+11=31.
*/
package main
import (
	"fmt"
	"os"
	"io"
)

func main() {
	var d int
	fmt.Fscan(os.Stdin, &d)
	var sum int

	for {
		_, err := fmt.Fscan(os.Stdin, &d)
		if err == io.EOF {
			break
		}

		sum += d
	}

	fmt.Printf("%d", sum)
}