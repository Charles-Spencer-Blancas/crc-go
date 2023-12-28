package main

import (
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

func doDivision(gen uint64, div uint64) uint64 {
	numBitsGen := numBits(gen)
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

	return work
}

func crc(gen uint64, data uint64) uint64 {
	numBitsGen := numBits(gen)
	div := data << (numBitsGen - 1)
	return (data << (numBitsGen - 1)) + doDivision(gen, div)
}

func checkCrc(gen uint64, msg uint64) bool {
	return doDivision(gen, msg) == 0
}
