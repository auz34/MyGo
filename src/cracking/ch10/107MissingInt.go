// Given an input file with four billion non-negative integers, provide an algorithm to generate an integer that
// is contained in the file. Assume you have 1GB of memory available for this task.
// FOLLOW UP
// What if you have only 10 MB of memory? Assume that all the values are distinct and we now have no more than one
// billion non-negative integers
package findMissingInt

func setBit(num, bitNumber uint) uint {
	return num | 1 << bitNumber
}

func getBit(num, bitNumber uint) bool {
	return (num & 1 << bitNumber) != 0
}

func searchInBucket(ar []uint, left, right uint) uint {
	size := (right - left) / 32 + 1
	mem := make([]uint, size)

	for i:=0; i<len(ar); i++ {
		if ar[i]<left || ar[i]>right {
			continue
		}

		n := ar[i] / 32
		b := ar[i] % 32
		mem[n] = setBit(mem[n], b)
	}

	for i:=left; i<right; i++ {
		n := ar[i] / 32
		b := ar[i] % 32

		if !getBit(mem[n], b) {
			return i
		}
	}

	return 0
}

func findMissing(ar []uint, memory uint) uint {
	bucketSize := memory * 8

	for l, r:=0, bucketSize; l<^uint(0); {
		candidate := searchInBucket(ar, l, r)
		if candidate != 0 {
			return candidate
		}
		l+=bucketSize
		r+=bucketSize
	}

	return 0
}