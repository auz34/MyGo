/*Sherlock Holmes suspects his archenemy, Professor Moriarty, is once again plotting something diabolical.
Sherlock's companion, Dr. Watson, suggests Moriarty may be responsible for MI6's recent issues with their supercomputer,
The Beast.

Shortly after resolving to investigate, Sherlock receives a note from Moriarty boasting about infecting The Beast
with a virus; however, he also gives him a clue—a number, N. Sherlock determines the key to removing the virus is
to find the largest Decent Number having N digits.

A Decent Number has the following properties:

Its digits can only be 3's and/or 5's.
The number of 3's it contains is divisible by 5.
The number of 5's it contains is divisible by 3.
If there are more than one such number, we pick the largest one.
Moriarty's virus shows a clock counting down to The Beast's destruction, and time is running out fast.
Your task is to help Sherlock find the key before The Beast is destroyed!

Constraints
1≤T≤20
1≤N≤100000

Input Format

The first line is an integer, T, denoting the number of test cases.

The T subsequent lines each contain an integer, N, detailing the number of digits in the number.

Output Format

Print the largest Decent Number having N digits; if no such number exists, tell Sherlock by printing -1.

Sample Input
4
1
3
5
11
Sample Output
-1
555
33333
55555533333
Explanation

For N=1, there is no decent number having 1 digit (so we print −1).
For N=3, 555 is the only possible number. The number 5 appears three times in this number, so our count of 5's is
evenly divisible by 3 (Decent Number Property 3).
For N=5, 33333 is the only possible number. The number 3 appears five times in this number, so our count of 3's is
evenly divisible by 5 (Decent Number Property 2).
For N=11, 55555533333 and all permutations of these digits are valid numbers; among them, the given number is
the largest one*/

package main

import (
	"fmt"
	"os"
)

func main() {
	var t int
	fmt.Fscan(os.Stdin, &t)
	for i:=0; i<t; i++ {
		var n int
		fmt.Fscan(os.Stdin, &n)
		solve2(n)
	}
}

func print2(n5, n3 int) {
	if n5 == -1 || n3 == -1 {
		fmt.Print("-1\n")
		return
	}

	for i:=0; i<n5; i++ {
		fmt.Print("5")
	}

	for i:=0; i<n3; i++ {
		fmt.Print("3")
	}

	fmt.Print("\n")
}

func solve2(n int) {
	rem := n % 3

	if rem == 0 {
		print2(n, 0)
		return
	}

	for c5 := n - rem; c5 >= 0; c5 -= 3 {
		if (n - c5) % 5 == 0 {
			print2(c5, n - c5)
			return
		}
	}

	print2(-1, -1)
}

