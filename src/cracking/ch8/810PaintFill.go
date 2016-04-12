// Implement the “paint  fill” function that one might see on many image editing programs. That is,
// given a screen (represented by a 2 dimensional array of Colors), a point, and a new color,  fill in the
// surrounding area until you hit a border of that color.’

package main

import "fmt"

const (
	R = 1
	G = 2
	B = 3
	W = 255
)


func main() {
	picture := make([][]int, 8)
	picture[0] = []int{W, W, W, W, W, W, W, W, W, W}
	picture[1] = []int{W, W, G, W, W, W, W, W, W, W}
	picture[2] = []int{W, G, G, G, G, G, W, W, W, W}
	picture[3] = []int{W, G, G, G, G, W, W, W, W, W}
	picture[4] = []int{W, G, G, G, G, W, W, W, W, W}
	picture[5] = []int{W, W, G, G, W, W, W, W, W, W}
	picture[6] = []int{W, W, W, W, W, W, W, W, W, W}
	picture[7] = []int{W, W, W, W, W, W, W, W, W, W}

	fmt.Print("Original picture: \n")
	printPicture(picture)

	fill(picture, 3, 3, B)
	fmt.Print("Filled picture: \n")
	printPicture(picture)

	fill(picture, 0, 0, R)
	fmt.Print("Changed background: \n")
	printPicture(picture)
}

func fill(picture [][]int, r, c, clr int) {
	originalClr := picture[r][c]

	if clr == originalClr {
		return
	}

	h := len(picture)
	w := len(picture[0])
	var paintPixel func(r, c int)
	paintPixel = func(r, c int) {
		picture[r][c] = clr
		if r>0 && picture[r-1][c] == originalClr {
			paintPixel(r-1, c)
		}

		if r<h-1 && picture[r+1][c] == originalClr {
			paintPixel(r+1, c)
		}

		if c>0 && picture[r][c-1] == originalClr {
			paintPixel(r, c-1)
		}

		if c<w-1 && picture[r][c+1] == originalClr {
			paintPixel(r, c+1)
		}
	}

	paintPixel(r, c)
}

func printPicture(picture [][]int) {
	for _, row := range picture {
		for _, pixel := range row {
			switch pixel {
			case R: fmt.Print("R ")
			case G: fmt.Print("G ")
			case B: fmt.Print("B ")
			case W: fmt.Print("W ")
			}
		}

		fmt.Print("\n")
	}
}
