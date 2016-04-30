/*
The Tower of Hanoi is a famous game consisting of 3 rods and a number of discs of incrementally different diameters.
The puzzle starts with the discs neatly stacked on one rod, ordered by ascending size with the smallest disc at the top.
The game's objective is to move the entire stack to another rod, obeying the following rules:

1) Only one disc can be moved at a time.
2) Each move consists of taking the topmost disc from a stack and moving it to the top of another stack.
3) No disc may be placed on top of a smaller disc.

Gena has a modified version of the Tower of Hanoi. His Hanoi has 4 rods and N discs ordered by ascending size. He made
a few moves (following the rules above), but stopped and lost his place. He wants to restore the tower to its original
state by making valid moves. Given the state of Gena's Hanoi, help him calculate the minimum number of moves needed to
restore the tower to its original state.

Note: Gena's rods are numbered from 1 to 4. All discs are initially located on rod 1.

Input Format
The first line contains a single integer, N, denoting the number of discs.
The second line contains N space-separated integers, where the ith integer is the index of the rod where the disk
with diameter i is located.

Constraints
1≤N≤10
10

Output Format
Print the minimum number of moves Gena must make to restore the tower to its initial, ordered state on the first rod.

Sample Input
3
1 4 1
Sample Output
3
Explanation
3 moves are enough to build the tower. Here is one possible solution
 <images>*/
package main

import (
	"fmt"
	"errors"
	"math"
)

type Stack struct {
	internal []int
}

func NewStack() *Stack {
	return &Stack{ make([]int, 0)}
}

func (self *Stack) Count() int {
	return len(self.internal)
}

func (self *Stack) Push(v int) {
	self.internal = append(self.internal, v)
}

func (self *Stack) Pop() (int, error) {
	l := len(self.internal)

	if l == 0 {
		return 0, errors.New("Empty stack")
	}

	res := self.internal[l-1]
	self.internal = self.internal[:l-1]

	return res, nil
}

func (self *Stack) Peek() (int, error) {
	l := len(self.internal)

	if l == 0 {
		return 0, errors.New("Empty stack")
	}

	res := self.internal[l-1]
	return res, nil
}

type Game struct {
	n int
	rods []*Stack
	curState map[int]int
	goalState map[int]int
}

func NewGame(n int, goal map[int]int) *Game {
	result := &Game{n, []*Stack { NewStack(), NewStack(), NewStack(), NewStack()}, make(map[int]int), goal}

	for i:=n; i>0; i-- {
		result.rods[0].Push(i)
		result.curState[i] = 0
	}

	return result
}

func (self *Game) move(from, to int) {
	disk, _ := self.rods[from].Pop()
	self.rods[to].Push(disk)
	self.curState[disk] = to
	fmt.Printf("Moved disk #%v from %v to %v \n", disk, from, to)
}

func Round(val float64, roundOn float64, places int ) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func magicK(n int) int {
	mid := Round(math.Sqrt(float64(2*n + 1)), .5, 3)
	return n - int(mid) + 1
}

func pickProxy(source, target int) (int, int) {
	mp := make(map[int]bool)
	proxy, left := -1, -1
	mp[source] = true
	mp[target] = true

	for i:=0; i<4; i++ {
		if !mp[i] {
			if proxy == -1 {
				proxy = i
			} else {
				left = i
			}
		}
	}

	return proxy, left
}

func (self *Game) hanoi(numOfDisks, source, target, temp int) {
	if numOfDisks == 1 {
		self.move(source, target)
		return
	}

	self.hanoi(numOfDisks-1, source, temp, target)
	self.move(source, target)
	self.hanoi(numOfDisks-1, temp, target, source)
}

func (self *Game) frameStuart(numOfDisks, source, target int) {
	k := magicK(numOfDisks)

	if numOfDisks == 1 {
		self.move(source, target)
		return
	}

	proxy, left := pickProxy(source, target)
	self.frameStuart(k, source, proxy)
	self.hanoi(numOfDisks - k, source, target, left)
	self.frameStuart(k, proxy, target)
}

func (self *Game) solve() {
	for i:=self.n; i>0; i-- {
		if self.curState[i] == self.goalState[i] {
			continue
		}

		proxy, _ := pickProxy(self.curState[i], self.goalState[i])
		self.frameStuart(i-1, self.curState[i], proxy)
		self.move(self.curState[i], self.goalState[i])
	}
}

func main() {
	ar := []int {1, 4, 1}
	goal := make(map[int]int)
	for i:=0; i<len(ar); i++ {
		goal[i+1]=ar[i]-1
	}

	game := NewGame(3, goal)
	game.solve()
}