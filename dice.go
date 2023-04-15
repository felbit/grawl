package main

import (
	"crypto/rand"
	"math/big"
)

func GetRandomInt(num int) int {
	return RollDice(num)
}

func GetRandomBetween(low, high int) int {
	var randy int = -1
	for {
		randy = RollDice(high)
		if randy >= low {
			break
		}
	}

	return randy
}

func RollDice(num int) int {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(num)))
	return int(x.Int64())
}
