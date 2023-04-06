package sl

import (
	"crypto/rand"
	"log"
	"math/big"
	"runtime/debug"
)

func rangeUint64(min uint64, max uint64) (random uint64) {
	// Swap min and max if min > max
	if min > max {
		min, max = max, min
		log.Print("goclub/slice: rangeUint64(min, max) max should greater than or equal to min", "\n", string(debug.Stack()))
	}

	// Calculate the range
	bigMax := new(big.Int).SetUint64(max)
	bigMin := new(big.Int).SetUint64(min)
	bigRange := new(big.Int).Sub(bigMax, bigMin)

	// Generate a random number within the range
	bigRandom, err := rand.Int(rand.Reader, bigRange.Add(bigRange, big.NewInt(1)))
	if err != nil {
		// 理论上不会出现这个错误,如果出错panic .
		// 这样既能保证函数易用,又能保证程序不会出现意外的错误
		panic(err)
	}
	random = bigRandom.Add(bigRandom, bigMin).Uint64()
	return random
}
