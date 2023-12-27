package main

import (
	"fmt"
	"math"
)

func numBits(x uint64) uint64 {
	y := float64(x)
	if y == 0 {
		y = 1
	}
	return 1 + uint64(math.Floor(math.Log2(y)))
}

func twoToPow(x uint64) uint64 {
	return uint64(math.Pow(2, float64(x)))
}

func crc(gen uint64, data uint64) uint64 {
	numBitsGen := numBits(gen)
	div := data << (numBitsGen - 1)

	numBitsDiv := numBits(div)

	var work uint64 = 0

	for i := uint64(0); i <= numBitsDiv; i++ {
		j := numBitsDiv - i

		if numBits(work) == numBitsGen {
			work = work ^ gen
		}

		pow2 := twoToPow(j)
		var toAdd uint64 = 0

		if div >= pow2 {
			div -= pow2
			toAdd = 1
		}

		work = work << 1
		work += toAdd
	}

	return (data << (numBitsGen - 1)) + work
}

func main() {
	var gen uint64 = 0b1011
	var data uint64 = 0b11011110

	fmt.Printf("%b\n", crc(gen, data))
}
