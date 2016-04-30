// implementation of Frame-Stewart algorithm for Hanoi towers with four pegs
package main

import (
	"fmt"
	"misc/math/structs"
	"math"
)

type Game struct {
	n int // number of disks
	pegs []*structs.Stack
}

func NewGame(n int) *Game {
	result := &Game{n, []*structs.Stack{ structs.NewStack(), structs.NewStack(), structs.NewStack(), structs.NewStack()}}

	for i:=n; i>0; i-- {
		result.pegs[0].Push(i)
	}

	return result
}

func (self *Game) move(from, to int) {
	if self.pegs[from].Count() == 0 {
		panic("Nothing to move")
	}

	disk, _ := self.pegs[from].Pop()

	if self.pegs[to].Count() > 0 {
		peek, _ := self.pegs[to].Peek()
		if peek < disk {
			panic("Break game rule")
		}
	}

	self.pegs[to].Push(disk)
	fmt.Printf("move disk #%v from %v to %v\n", disk, from, to)
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
	self.frameStuart(self.n, 0, 3)
}

func main() {
	game := NewGame(5)
	game.solve()
}


