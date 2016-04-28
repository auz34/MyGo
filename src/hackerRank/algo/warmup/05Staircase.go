/* Your teacher has given you the task of drawing a staircase structure. Being an expert programmer,
you decided to make a program to draw it for you instead. Given the required height, can you print
a staircase as shown in the example?

Input
You are given an integer N depicting the height of the staircase.

Output
Print a staircase of height N that consists of # symbols and spaces. For example for N=6, here's
a staircase of that height:

     #
    ##
   ###
  ####
 #####
######
Note: The last line has 0 spaces before it.*/

package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Fscan(os.Stdin, &n)
	for i := 1; i<=n; i++ {
		printLine(i, n)
	}
}

func printLine(step, height int) {
	for i:=height; i>0; i-- {
		if i>step {
			fmt.Print(" ")
		} else {
			fmt.Print("#")
		}
	}

	fmt.Print("\n")
}