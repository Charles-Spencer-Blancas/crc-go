package main

import (
	"errors"
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

func doDivision(gen uint64, div uint64) (uint64, error) {
	if gen <= 1 {
		return 0, errors.New("generator must be > 1")
	}
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

	// If in the last iteration we can still divide by gen, do iteration
	for numBits(work) >= numBitsGen {
		work = work ^ gen
	}

	return work, nil
}

func crc(gen uint64, data uint64) (uint64, error) {
	numBitsGen := numBits(gen)
	div := data << (numBitsGen - 1)
	rem, err := doDivision(gen, div)
	if err != nil {
		return data, err
	}
	return (data << (numBitsGen - 1)) + rem, nil
}

func checkCrc(gen uint64, msg uint64) (bool, error) {
	answer, err := doDivision(gen, msg)
	if err != nil {
		return false, err
	}
	return answer == 0, nil
}

func decodeCrc(gen uint64, msg uint64) (uint64, error) {
	valid, err := checkCrc(gen, msg)
	if err != nil {
		return msg, err
	}

	if valid {
		return msg >> (numBits(gen) - 1), nil
	} else {
		return msg, errors.New("there is an error in the message")
	}
}
